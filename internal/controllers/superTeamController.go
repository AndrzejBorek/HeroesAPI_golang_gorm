package controllers

import (
	"fmt"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/requestModels"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strconv"
)

func GetSuperTeams(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var superTeams []models.SuperTeam
		err := db.Find(&superTeams).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
		}
		c.JSON(http.StatusOK, superTeams)
	}
}

func GetMembersOfSuperTeam(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var heroes []models.Hero
		teamID, err := strconv.ParseUint(c.Param("teamID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong teamID format."})
			return
		}

		err = db.First(&models.SuperTeam{}, teamID).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "SuperTeam with given ID does not exist."})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
			return
		}

		err = db.Where("super_team_id = ?", teamID).Find(&heroes).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
			return
		}
		if len(heroes) == 0 {
			c.JSON(http.StatusNoContent, gin.H{"error": "This SuperTeam does not have any heroes."})
			return
		}
		c.JSON(http.StatusOK, heroes)
	}
}

func CreateSuperTeam(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var superTeamInput requestModels.CreateSuperTeamInput
		err := c.ShouldBindJSON(&superTeamInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if superTeamInput.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty."})
			return
		}
		superTeamName := superTeamInput.Name

		var heroes []*models.Hero
		err = db.Find(&heroes, superTeamInput.HeroesIDs).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(heroes) != len(superTeamInput.HeroesIDs) {
			missingIDs := make([]uint, 0)
			existingIDs := make(map[uint]bool)
			for _, sp := range heroes {
				existingIDs[sp.ID] = true
			}
			for _, id := range superTeamInput.HeroesIDs {
				if !existingIDs[id] {
					missingIDs = append(missingIDs, id)
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"Code": "Not Found", "Error": "HeroesIDs with given IDs do not exist.", "HeroesIDs": missingIDs})
			return
		}

		superTeam := models.SuperTeam{
			Name:   superTeamName,
			Heroes: heroes,
		}

		result := db.Create(&superTeam)

		if result.Error != nil {
			fmt.Print(reflect.TypeOf(result.Error))
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}

		c.JSON(http.StatusCreated, &superTeam)
	}
}
