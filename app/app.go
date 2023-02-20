package app

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/controllers"
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/database"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter() *gin.Engine {
	db, err := database.GenerateDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	r := gin.Default()

	r.GET("/heroes", controllers.GetHeroes(db))
	r.GET("/heroes/:heroID/helpers", controllers.GetHeroHelpers(db))
	r.DELETE("/heroes/:heroID", controllers.DeleteHero(db))
	r.POST("/heroes", controllers.CreateHero(db))

	r.GET("/superpowers", controllers.GetSuperPowers(db))
	r.POST("/superpowers", controllers.AddSuperPower(db))

	r.GET("/helpers", controllers.GetHelpers(db))

	r.GET("/villains", controllers.GetVillains(db))
	r.POST("/villains", controllers.CreateVillain(db))
	r.DELETE("/villains/:villainID", controllers.DeleteVillain(db))
	return r
}
