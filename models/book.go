package models

type Book struct {
	BaseModel
	Name          string `gorm:"type:string;size:20;not null"`
	NumberOfPages int    `gorm:"type:int;not null"`
	Description   string `gorm:"type:string;size:250"`
	Author        Author `gorm:"foreignKey:AuthorId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	AuthorId      int
}
