package main

import (
	"log"
	"refresh_token/config"
	"refresh_token/internal/database"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal("error loading env: ", err)
	}

	db, err := config.Setup()
	if err != nil {
		log.Fatal("error setting up database: ", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal("error migrating database: ", err)
	}
}
