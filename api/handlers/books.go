package handlers

import (
	"library-web-api-go/config"
	"library-web-api-go/services"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(cfg *config.Config) *BookHandler {
	return &BookHandler{
		service: services.NewBookService(cfg),
	}
}

func (h *BookHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *BookHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *BookHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *BookHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}
