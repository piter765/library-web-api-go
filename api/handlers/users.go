package handlers

import (
	"library-web-api-go/api/dto"
	"library-web-api-go/api/helpers"
	"library-web-api-go/config"
	"library-web-api-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{service: service}
}

func (h *UsersHandler) SignIn(c *gin.Context) {
	req := new(dto.SignInRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, helpers.ValidationError, err))
		return
	}
	tokenDetail, err := h.service.SignIn(req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(tokenDetail, true, helpers.Success))
}
