package routers

import (
	"library-web-api-go/api/handlers"
	"library-web-api-go/config"

	"github.com/gin-gonic/gin"
)

func Author(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewAuthorHandler(cfg)

	router.POST("/", h.Create)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
	router.GET("/:id", h.GetById)
}