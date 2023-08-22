package database

import (
	"errors"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func GenerateDatabase() (*gorm.DB, error) {
	var err error
	Db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	// TODO check how this option in gorm.Config will affect performance of app: , SkipDefaultTransaction: true
	if err != nil {
		return nil, err
	}
	err = Db.AutoMigrate(&models.SuperPower{}, &models.Hero{}, &models.Helper{}, &models.Villain{}, &models.SuperTeam{}, &models.EvilPlan{})
	if err != nil {
		return nil, err
	}
	err = FillDatabase(Db)
	if err != nil {
		log.Fatalf("Failed to insert data into database: %v", err)
		return nil, err
	}
	return Db, nil
}

func FillDatabase(db *gorm.DB) error {

	superStrength := models.SuperPower{
		Description: "Strength of ten people.",
	}
	result := db.Create(&superStrength)
	if result.Error != nil {
		return errors.New("error while inserting superStrength")
	}

	hero := models.Hero{Name: "Superman"}
	result = db.Create(&hero)
	if result.Error != nil {
		return errors.New("error while inserting hero")
	}

	if err := db.Model(&hero).Association("SuperPowers").Append(&superStrength); err != nil {
		return errors.New("error while associating hero and superpower")
	}

	helper := models.Helper{Name: "Helper1", HeroID: hero.ID}
	result = db.Create(&helper)
	if result.Error != nil {
		return errors.New("error while inserting helper1")
	}

	helper2 := models.Helper{Name: "Helper2", HeroID: hero.ID}
	result = db.Create(&helper2)
	if result.Error != nil {
		return errors.New("error while inserting helper2")
	}

	return nil
}
