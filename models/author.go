package models

import "time"

type Author struct {
	FirstName   string    `gorm:"type:string;size:20;not null"`
	LastName    string    `gorm:"type:string;size:20;not null"`
	DateOfBirth time.Time `gorm:"type:TIMESTAMP;not null"`
	Books       *[]Book
}