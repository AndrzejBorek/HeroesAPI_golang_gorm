package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	ID         uint         `json:"ID" gorm:"primary_key"`
	Name       string       `json:"Name" gorm:"Not Null;unique"`
	SuperPower []SuperPower `json:"SuperPowers" gorm:"unique;many2many:hero_superpowers;association_foreignkey:ID;foreignkey:ID"`
	Villains   []Villain    `json:"Villains" gorm:"many2many:hero_villains;association_foreignkey:ID;foreignkey:ID"`
	Helper     *Helper      `json:"Helper" gorm:"foreignKey:Hero;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
