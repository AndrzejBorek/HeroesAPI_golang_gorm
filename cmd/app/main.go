package main

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/app"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	db, err := database.GenerateDatabase(os.Getenv("DATABASE_DSN"), true)
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	r := app.SetupRouter(db)
	err = r.Run()
	if err != nil {
		return
	}
}
