package main

import (
	api "library-web-api-go"
	"library-web-api-go/config"
	"library-web-api-go/database"
	"library-web-api-go/database/migrations"
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
	migrations.Up_1()

	api.InitServer(cfg)
}
