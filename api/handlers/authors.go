package handlers

import (
	"library-web-api-go/config"
	"library-web-api-go/services"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	service *services.AuthorService
}

func NewAuthorHandler(cfg *config.Config) *AuthorHandler {
	return &AuthorHandler{
		service: services.NewAuthorService(cfg),
	}
}

func (h *AuthorHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *AuthorHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *AuthorHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *AuthorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}
