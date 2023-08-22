package main

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/app"
	"log"
)

func main() {
	db, err := database.GenerateDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	r := app.SetupRouter(db)
	err = r.Run()
	if err != nil {
		return
	}
}
