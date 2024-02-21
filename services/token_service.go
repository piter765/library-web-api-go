package services

import (
	"library-web-api-go/api/dto"
	"library-web-api-go/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	cfg *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	Email        string
	MobileNumber string
	//roles
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s *TokenService) GenerateTokens(token *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[UserIdKey] = token.UserId
	atc[FirstNameKey] = token.FirstName
	atc[LastNameKey] = token.LastName
	atc[UsernameKey] = token.Username
	atc[EmailKey] = token.Email
	atc[MobileNumberKey] = token.MobileNumber
	atc[ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[UserIdKey] = token.UserId
	rtc[ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

const (
	// Claims
	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "UserId"
	FirstNameKey           string = "FirstName"
	LastNameKey            string = "LastName"
	UsernameKey            string = "Username"
	EmailKey               string = "Email"
	MobileNumberKey        string = "MobileNumber"
	ExpireTimeKey          string = "Exp"
)
