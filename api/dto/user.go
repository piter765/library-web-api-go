package dto

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

type SignInRequest struct {
	Email    string
	Password string
}

type SignUpRequest struct {
	Email string
	FirstName string
	LastName string
	Password string
}