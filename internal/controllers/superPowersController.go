package controllers

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database/queries"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/requestModels"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AddSuperPower(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var superPowerInput requestModels.CreateSuperPowerInput
		err := c.ShouldBindJSON(&superPowerInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Description is missing."})
			return
		}

		superPower := models.SuperPower{
			Description: superPowerInput.Description,
		}
		result := db.Create(&superPower)
		err = result.Error
		//TODO: When gorm will implement error for failed unique constraint check if this error occurs and return 409 code
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"SuperPower": superPower})
	}
}

func GetSuperPowers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var superPowers []models.SuperPower
		superPowers, err := queries.GetAllSuperPowers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
		}
		c.JSON(http.StatusOK, superPowers)
	}
}
