package api

import (
	"fmt"
	"library-web-api-go/config"
	"log"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	//middlewares


	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		log.Fatal("Error running the server")
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
}
