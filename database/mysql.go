package database

import (
	"fmt"
	"library-web-api-go/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.MySql.User, cfg.MySql.Password, cfg.MySql.Host, cfg.MySql.Port, cfg.MySql.DbName)

	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			return err
	}

	sqlDb, err := dbClient.DB()
	if err != nil {
			return err
	}

	err = sqlDb.Ping()
	if err != nil {
			return err
	}

	sqlDb.SetMaxIdleConns(cfg.MySql.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.MySql.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.MySql.ConnMaxLifetime * time.Minute)

	log.Println("DB connection established")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}