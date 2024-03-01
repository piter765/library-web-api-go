package migrations

import (
	"library-web-api-go/consts"
	"library-web-api-go/database"
	"library-web-api-go/models"
	"time"

	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Up_1() {
	database := database.GetDb()

	createTables(database)
	createDefaultUserInformation(database)
	createAuthorWithBooks(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	tables = addNewTable(database, models.Author{}, tables)

	tables = addNewTable(database, models.Book{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		fmt.Println("Error creating the tables in database")
	}
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {
	adminRole := models.Role{Name: consts.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: consts.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: consts.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "123456789", Email: "admin@admin.com"}

	password := "test1234"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("email = ?", u.Email).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createAuthorWithBooks(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Author{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Author{
			FirstName:   "George",
			LastName:    "Martin",
			DateOfBirth: time.Date(1960, time.January, 2, 15, 4, 5, 0, time.Local), //should change the date
			Books: []*models.Book{
				{
					Name:          "Book 1",
					NumberOfPages: 100,
					Description:   "Description for Book 1",
				},
				{
					Name:          "Book 2",
					NumberOfPages: 150,
					Description:   "Description for Book 2",
				},
			},
		})
	}
}
