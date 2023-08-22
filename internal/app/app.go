package app

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	//db, err := database.GenerateDatabase()
	//if err != nil {
	//	log.Fatalf("Failed to setup database: %v", err)
	//}

	r := gin.Default()
	//HeroesIDs
	r.GET("/heroes", controllers.GetHeroes(db))
	r.GET("/heroes/:heroID/helpers", controllers.GetHeroHelpers(db))
	r.GET("/heroes/:heroID", controllers.GetHeroByID(db))
	r.DELETE("/heroes/:heroID", controllers.DeleteHero(db))
	r.POST("/heroes", controllers.CreateHero(db))

	//Superpowers
	r.GET("/superpowers", controllers.GetSuperPowers(db))
	r.POST("/superpowers", controllers.AddSuperPower(db))

	//Helpers
	r.GET("/helpers", controllers.GetHelpers(db))

	// Villains
	r.GET("/villains", controllers.GetVillains(db))
	r.POST("/villains", controllers.CreateVillain(db))
	r.DELETE("/villains/:villainID", controllers.DeleteVillain(db))
	//EvilPlan

	// Super teams
	r.GET("/superteam/", controllers.GetSuperTeams(db))
	r.GET("/superteam/:teamID", controllers.GetMembersOfSuperTeam(db))
	r.POST("/superteam/", controllers.CreateSuperTeam(db))

	return r
}
