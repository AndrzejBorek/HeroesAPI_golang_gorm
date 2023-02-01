package models

import "gorm.io/gorm"

type Villain struct {
	gorm.Model
	ID           uint         `json:"ID" gorm:"primary_key"`
	Name         string       `json:"name"`
	VillainPower []SuperPower `json:"superPowers" gorm:"many2many:villain_villainpowers;"`
	HeroEnemy    []Hero       `json:"heroEnemy" gorm:"many2many:hero_villains;association_foreignkey:ID;foreignkey:ID"`
}
