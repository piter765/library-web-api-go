package api

import (
	"fmt"
	"library-web-api-go/api/routers"
	"library-web-api-go/config"
	"log"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	//middlewares

	RegisterRoutes(r, cfg)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		log.Fatal("Error running the server")
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("v1")
	{
		//User 
		users := v1.Group("/users")

		routers.User(users, cfg)
	}
}
