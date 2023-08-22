package controllers

import (
	"fmt"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database/queries"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/requestModels"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetHeroes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var heroes []models.Hero
		heroes, err := queries.GetHeroes(db)
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
		heroID, err := utils.StringToUint(c.Param("heroID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong ID format."})
			return
		}
		err = queries.DeleteHeroById(db, heroID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Hero with ID: " + c.Param("heroID") + " not found."})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Writer.WriteHeader(http.StatusNoContent)
	}
}

func GetHeroByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var hero models.Hero
		heroID, err := utils.StringToUint(c.Param("heroID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong ID."})
			return
		}
		hero, err = queries.GetHeroById(db, heroID)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Hero with ID: " + c.Param("heroID") + " not found."})
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
		hero, err := queries.CreateHero(db, heroInput.Name, heroInput.SuperPowersIDs, heroInput.VillainsIDs, heroInput.SuperTeamsIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//TODO: When gorm will implement error for failed unique constraint check if this error occurs and return 409 code

		c.JSON(http.StatusCreated, hero)
	}
}
