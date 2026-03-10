package main

import (
	"log"
	"refresh_token/internal/app"
)

func main() {
	db := SetupDatabase()

	a := app.New(db)
	err := a.Start()
	if err != nil {
		log.Fatal("error starting app: ", err)
	}
}
