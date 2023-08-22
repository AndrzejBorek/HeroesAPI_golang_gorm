package controllers

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database/queries"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetHelpers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var helpers []models.Helper
		helpers, err := queries.GetAllHelpers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during database query."})
		}
		c.JSON(http.StatusOK, helpers)
	}
}
