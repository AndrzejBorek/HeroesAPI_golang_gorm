package controllers

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/models"
	"gorm.io/gorm"
)

func GetHeroes(db *gorm.DB) []models.Hero {
	var heroes []models.Hero
	db.Preload("SuperPower").Find(&heroes)
	return heroes
}
