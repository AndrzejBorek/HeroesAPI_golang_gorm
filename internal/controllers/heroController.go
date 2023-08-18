package controllers

import (
	"fmt"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/requestModels"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetHeroes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var heroes []models.Hero
		err := db.Model(&models.Hero{}).Preload("Villains").Preload("Helpers").Preload("SuperPowers").Preload("SuperTeams").Find(&heroes).Error
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
			return
		}
		c.JSON(http.StatusOK, heroes)
	}
}

func GetHeroHelpers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var helpers []*models.Helper
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

func GetHeroByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var hero models.Hero
		heroID := c.Param("heroID")
		err := db.Preload("Helpers").Preload("Villains").Preload("SuperPowers").Preload("SuperTeams").First(&hero, heroID).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Hero with ID: " + heroID + " not found."})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, hero)
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

		var helpers []*models.Helper
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

		var superTeams []*models.SuperTeam
		if len(heroInput.SuperTeamsIDs) > 0 {
			result = db.Find(&superTeams, heroInput.SuperTeamsIDs)
			if len(superTeams) != len(heroInput.SuperTeamsIDs) {
				missingIDs := make([]uint, 0)
				existingIDs := make(map[uint]bool)
				for _, sp := range superTeams {
					existingIDs[sp.ID] = true
				}
				for _, id := range heroInput.SuperTeamsIDs {
					if !existingIDs[id] {
						missingIDs = append(missingIDs, id)
					}
				}
				c.JSON(http.StatusNotFound, gin.H{"Code": "Not Found", "Error": "SuperTeams with given IDs do not exist.", "SuperTeams": missingIDs})
				return
			}
		}
		hero := models.Hero{
			Name:        heroInput.Name,
			SuperPowers: superPowers,
			Villains:    villains,
			Helpers:     helpers,
			SuperTeams:  superTeams,
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
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
