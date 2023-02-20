package controllers

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/requestModels"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetHeroes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var heroes []models.Hero
		err := db.Preload("Villains").Preload("Helpers").Preload("SuperPowers").Find(&heroes).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
		}
		c.JSON(http.StatusOK, heroes)
	}
}

func GetHeroHelpers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var helpers []models.Helper
		heroID := c.Param("heroID")
		err := db.First(&models.Hero{}, heroID).Error
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": "Not Found", "error": "Hero with ID: " + heroID + " not found."})
			return
		}
		err = db.Where("hero_ID = ?", heroID).Find(&helpers).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
			return
		}
		c.JSON(http.StatusOK, helpers)
	}
}

func DeleteHero(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var hero models.Hero
		heroID := c.Param("heroID")
		err := db.First(&hero, heroID).Error
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hero with ID: " + heroID + " not found."})
			return
		}
		result := db.Select("Helpers").Delete(&hero)
		err = result.Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}

func CreateHero(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var heroInput requestModels.CreateHeroInput
		err := c.ShouldBindJSON(&heroInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var superPowers []models.SuperPower
		result := db.Find(&superPowers, heroInput.SuperPowersIDs)
		if len(superPowers) != len(heroInput.SuperPowersIDs) {
			missingIDs := make([]uint, 0)
			existingIDs := make(map[uint]bool)
			for _, sp := range superPowers {
				existingIDs[sp.ID] = true
			}
			for _, id := range heroInput.SuperPowersIDs {
				if !existingIDs[id] {
					missingIDs = append(missingIDs, id)
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"Code": "Not Found", "Error": "SuperPowers with given IDs do not exist.", "SuperPowers": missingIDs})
			return
		}
		// Second approach. I choose first one, because this had to query database n times, where n is len of superPowers
		// from request body. Although it had much easier checking if superPower does not exist.

		//for _, id := range heroInput.SuperPowersIDs {
		//	var superPower models.SuperPower
		//	if errors.Is(db.First(&superPower, id).Error, gorm.ErrRecordNotFound) {
		//		c.JSON(http.StatusNotFound, gin.H{"error": "SuperPower with ID: " + strconv.Itoa(int(id)) + " not found."})
		//		return
		//	}
		//	superPowers = append(superPowers, superPower)
		//}

		//var villains []*models.Villain
		//if heroInput.VillainsIDs != nil {
		//	for _, id := range heroInput.VillainsIDs {
		//		var villain models.Villain
		//		err := db.First(&villain, id).Error
		//		if err == gorm.ErrRecordNotFound {
		//			c.JSON(http.StatusBadRequest, gin.H{"error": "Villain with ID: " + strconv.Itoa(int(id)) + " not found."})
		//			return
		//		}
		//		villains = append(villains, &villain)
		//	}
		//}
		var villains []*models.Villain
		if len(heroInput.VillainsIDs) > 0 {
			result = db.Find(&villains, heroInput.VillainsIDs)
			if len(villains) != len(heroInput.VillainsIDs) {
				missingIDs := make([]uint, 0)
				existingIDs := make(map[uint]bool)
				for _, sp := range villains {
					existingIDs[sp.ID] = true
				}
				for _, id := range heroInput.VillainsIDs {
					if !existingIDs[id] {
						missingIDs = append(missingIDs, id)
					}
				}
				c.JSON(http.StatusNotFound, gin.H{"Code": "Not Found", "Error": "Villains with given IDs do not exist.", "Villains": missingIDs})
				return
			}
		}

		var helpers []models.Helper
		if len(heroInput.HelpersIDs) > 0 {
			result = db.Find(&helpers, heroInput.HelpersIDs)
			if len(helpers) != len(heroInput.HelpersIDs) {
				missingIDs := make([]uint, 0)
				existingIDs := make(map[uint]bool)
				for _, sp := range helpers {
					existingIDs[sp.ID] = true
				}
				for _, id := range heroInput.HelpersIDs {
					if !existingIDs[id] {
						missingIDs = append(missingIDs, id)
					}
				}
				c.JSON(http.StatusNotFound, gin.H{"Code": "Not Found", "Error": "Helpers with given IDs do not exist.", "Helpers": missingIDs})
				return
			}
		}
		//if heroInput.HelpersIDs != nil {
		//	for _, id := range heroInput.HelpersIDs {
		//		var helper models.Helper
		//		err := db.First(&helper, id).Error
		//		if err == gorm.ErrRecordNotFound {
		//			c.JSON(http.StatusBadRequest, gin.H{"error": "Helper with ID: " + strconv.Itoa(int(id)) + " not found."})
		//			return
		//		}
		//		helpers = append(helpers, helper)
		//	}
		//}

		hero := models.Hero{
			Name:        heroInput.Name,
			SuperPowers: superPowers,
			Villains:    villains,
			Helpers:     helpers,
		}

		result = db.Create(&hero)
		//TODO: When gorm will implement error for failed unique constraint check if this error occurs and return 409 code
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusCreated, &hero)
	}
}
