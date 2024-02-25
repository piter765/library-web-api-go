package services

import (
	"library-web-api-go/api/dto"
	"library-web-api-go/config"
	"library-web-api-go/database"
	"library-web-api-go/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	tokenService *TokenService
	database     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := database.GetDb()
  return &UserService{
		cfg: cfg,
		database: database,
		tokenService: NewTokenService(cfg),
	}
}

func (s *UserService) SignIn(req *dto.SignInRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.
		Model(&models.User{}).
		Where("email: ?", req.Email).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}
	
	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}

	tokenDetail, err := s.tokenService.GenerateTokens(&tdto)
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}
