package queries

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"gorm.io/gorm"
)

func GetAllHelpers(db *gorm.DB) ([]models.Helper, error) {
	var helpers []models.Helper
	err := db.Find(&helpers).Error
	if err != nil {
		return nil, err
	}
	return helpers, nil
}
