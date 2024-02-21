package main

import (
	api "library-web-api-go"
	"library-web-api-go/config"
	"library-web-api-go/database"
	"log"
)

func main() {

	cfg := config.GetConfig()

	err := database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		log.Fatal("Error running the database")
	}
	//migrations

	api.InitServer(cfg)
}
