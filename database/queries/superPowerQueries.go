package queries

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"gorm.io/gorm"
)

func CreateSuperPower(db *gorm.DB, description string) error {

	return nil
}

func GetAllSuperPowers(db *gorm.DB) ([]models.SuperPower, error) {
	var superPowers []models.SuperPower
	err := db.Find(&superPowers).Error
	if err != nil {
		return nil, err
	}
	return superPowers, nil
}
