package models

type User struct {
	BaseModel
	Email        string `gorm:"type:string;size:30;not null;unique"`
	FirstName    string `gorm:"type:string;size:20;not null"`
	LastName     string `gorm:"type:string;size:20;not null"`
	Password     string `gorm:"type:string;size:50;not null"`
	Username     string `gorm:"type:string;size:20;not null"`
	MobileNumber string `gorm:"type:string;size:20"`
	UserRoles    *[]UserRole
}

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:20;not null;unique"`
	UserRoles *[]UserRole
}

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gomr:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	RoleId int
}
