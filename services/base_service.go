package services

import (
	"context"
	"database/sql"
	"fmt"
	"library-web-api-go/common"
	"library-web-api-go/database"
	"library-web-api-go/models"
	"library-web-api-go/pkg/service_errors"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Preloads []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any]() *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: database.GetDb(),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := common.TypeConverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.
				  Create(model).
					Error
	if err != nil {
		tx.Rollback()
		fmt.Println("Cannot create ", reflect.TypeOf(*model).String())
		return nil, err
	}
	tx.Commit()
	bm, _ := common.TypeConverter[models.BaseModel](model)
	return s.GetById(bm.Id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)
	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v;
	}
	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	if err := tx.Model(model).
							Where("id = ? and deleted_at is null", id).
							Updates(snakeMap).
							Error; err != nil {
								tx.Rollback()
								fmt.Println("Cannot update ", reflect.TypeOf(*model).String())
								return nil, err
							}
	tx.Commit()
	return s.GetById(id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()
	model := new(T)

  deleteMap := map[string]interface{}{
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now()},
	}

	if count := tx.
								Model(model).
								Where("id = ? and deleted_at is null", id).
								Updates(deleteMap).
								RowsAffected; count == 0 {
									tx.Rollback()
									fmt.Println("Cannot delete ", reflect.TypeOf(*model).String())
									return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
								}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(id int) (*Tr, error) {
	model := new(T)
	db := Preload(s.Database, s.Preloads)
	err := db.Where("id = ? and deleted_at is null", id).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return common.TypeConverter[Tr](model)
}

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}
	return db
}