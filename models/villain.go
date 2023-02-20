package models

import "gorm.io/gorm"

type Villain struct {
	gorm.Model
	Name        string       `json:"name" gorm:"Not Null"`
	SuperPowers []SuperPower `json:"superPowers" gorm:"unique;many2many:villain_superpowers;"`
	HeroEnemies []*Hero      `json:"heroEnemies" gorm:"many2many:hero_villains"`
}

func (Villain) TableName() string {
	return "Villain"
}
