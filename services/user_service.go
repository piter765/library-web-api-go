package services

import (
	"library-web-api-go/config"

	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	tokenService *TokenService
	database     *gorm.DB
}
