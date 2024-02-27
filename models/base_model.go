package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP;null"`

	// CreatedBy  int            `gorm:"not null"`
	// ModifiedBy *sql.NullInt64 `gorm:"null"`
	// DeletedBy  *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return 
}
