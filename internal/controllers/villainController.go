package controllers

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/requestModels"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetVillains(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var villains []models.Villain
		db.Preload("HeroEnemies").Find(&villains)
		c.JSON(http.StatusOK, villains)
	}
}

func CreateVillain(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var villainInput requestModels.CreateVillainInput
		err := c.ShouldBindJSON(&villainInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var superPowers []*models.SuperPower
		for _, id := range villainInput.VillainPowersIDs {
			var superPower *models.SuperPower
			err := db.First(&superPower, id).Error
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, gin.H{"error": "SuperPower with ID: " + strconv.Itoa(int(id)) + " not found."})
				return
			}
			superPowers = append(superPowers, superPower)
		}
		var heroEnemies []*models.Hero
		for _, id := range villainInput.HeroEnemiesIDs {
			var hero models.Hero
			err := db.First(&hero, id).Error
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, gin.H{"error": "SuperHero with ID: " + strconv.Itoa(int(id)) + " not found."})
				return
			}
			heroEnemies = append(heroEnemies, &hero)
		}

		villain := models.Villain{
			Name:        villainInput.Name,
			SuperPowers: superPowers,
			HeroEnemies: heroEnemies,
		}
		result := db.Create(&villain)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusCreated, villain)
	}
}

func DeleteVillain(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var villain models.Villain
		villainID := c.Param("villainID")
		err := db.First(&villain, villainID).Error
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Villain with ID: " + villainID + " not found."})
			return
		}
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
