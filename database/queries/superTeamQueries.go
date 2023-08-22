package queries

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"gorm.io/gorm"
)

func GetSuperTeamMembers(db *gorm.DB, teamID uint) ([]models.HeroNameID, error) {

	var team models.SuperTeam
	err := db.Preload("Heroes").First(&team, teamID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	serializedHeroes := make([]models.HeroNameID, len(team.Heroes))
	for i, hero := range team.Heroes {
		serializedHeroes[i] = models.HeroNameID{
			ID:   hero.ID,
			Name: hero.Name,
		}
	}
	return serializedHeroes, nil

}
