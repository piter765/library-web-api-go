package routers

import (
	"library-web-api-go/api/handlers"
	"library-web-api-go/config"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("sign-in", h.SignIn)
}