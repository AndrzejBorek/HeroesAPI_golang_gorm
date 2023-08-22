package queries

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/utils"
	"gorm.io/gorm"
)

func GetHeroes(db *gorm.DB) ([]models.Hero, error) {
	var heroes []models.Hero
	err := db.Preload("Villains").Preload("Helpers").Preload("SuperPowers").Preload("SuperTeams").Find(&heroes).Error
	return heroes, err
}

func GetHeroById(db *gorm.DB, heroID uint) (models.Hero, error) {
	var hero models.Hero
	err := db.Preload("Villains").Preload("Helpers").Preload("SuperPowers").Preload("SuperTeams").First(&hero, heroID).Error
	if err != nil {
		return models.Hero{}, err
	}
	return hero, nil
}

func DeleteHeroById(db *gorm.DB, heroID uint) error {
	var hero models.Hero
	err := db.First(&hero, heroID).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	err = db.Unscoped().Where("hero_id = ?", heroID).Delete(&models.Helper{}).Error
	if err != nil {
		return err
	}
	err = db.Delete(&hero).Error
	if err != nil {
		return err
	}
	return nil

}

func CreateHero(db *gorm.DB, name string, superPowerIDs []uint, villainIDs []uint, superTeamsIDs []uint) (models.Hero, error) {
	var superPowers []models.SuperPower

	if len(superPowerIDs) > 0 {
		err := db.Find(&superPowers, superPowerIDs).Error
		if err != nil {
			return models.Hero{}, err
		}
		if len(superPowers) != len(superPowerIDs) {
			superPowersValidator := func(item interface{}) uint {
				return item.(models.SuperPower).ID
			}
			superPowersInterface, err := utils.SliceToInterface(superPowers)
			if err != nil {
				return models.Hero{}, err
			}
			err = utils.ValidateItemIDs(superPowersInterface, superPowerIDs, "SuperPowers with given IDs do not exist: ", superPowersValidator)
			if err != nil {
				return models.Hero{}, err
			}
		}
	}

	var villains []*models.Villain
	if len(villainIDs) > 0 {
		err := db.Find(&villains, villainIDs).Error
		if err != nil {
			return models.Hero{}, err
		}
		if len(villains) != len(villainIDs) {
			villainsValidator := func(item interface{}) uint {
				return item.(*models.Villain).ID
			}
			villainsInterface, err := utils.SliceToInterface(villains)
			if err != nil {
				return models.Hero{}, err
			}
			err = utils.ValidateItemIDs(villainsInterface, villainIDs, "Villains with given IDs do not exist: ", villainsValidator)
			if err != nil {
				return models.Hero{}, err
			}
		}
	}

	var superTeams []*models.SuperTeam
	if len(superTeamsIDs) > 0 {
		err := db.Find(&superTeams, superTeamsIDs).Error
		if err != nil {
			return models.Hero{}, err
		}
		if len(superTeams) != len(superTeamsIDs) {
			superTeamsValidator := func(item interface{}) uint {
				return item.(*models.SuperTeam).ID
			}
			superTeamsInterface, err := utils.SliceToInterface(superTeams)
			if err != nil {
				return models.Hero{}, err
			}
			err = utils.ValidateItemIDs(superTeamsInterface, superTeamsIDs, "SuperTeams with given IDs do not exist: ", superTeamsValidator)
			if err != nil {
				return models.Hero{}, err
			}
		}
	}

	hero := models.Hero{
		Name:        name,
		SuperPowers: superPowers,
		Villains:    villains,
		SuperTeams:  superTeams,
	}
	err := db.Create(&hero).Error

	if err != nil {
		return models.Hero{}, err
	}
	return hero, nil
}
