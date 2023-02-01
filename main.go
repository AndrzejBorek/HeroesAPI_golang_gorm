package main

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/controllers"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("HeroesAPI.db"), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.Migrator().DropTable(&models.Hero{}, &models.SuperPower{})
	if err != nil {
		return
	}
	dbErr := db.AutoMigrate(&models.SuperPower{}, &models.Hero{})
	if dbErr != nil {
		return
	}

	superStrength := models.SuperPower{
		Description: "Strength of ten people.",
	}

	db.Create(&superStrength)
	hero := models.Hero{Name: "Superman"}
	db.Create(&hero)

	err = db.Model(&hero).Association("SuperPower").Append(&superStrength)
	if err != nil {
		return
	}

	r := gin.Default()
	r.GET("/heroes", func(c *gin.Context) {
		c.JSON(http.StatusOK, controllers.GetHeroes(db))
	})
	err = r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080

}
