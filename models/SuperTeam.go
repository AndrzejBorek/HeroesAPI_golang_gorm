package models

import "gorm.io/gorm"

type SuperTeam struct {
	gorm.Model
	ID uint `json:"ID" gorm:"primary_key"`
	//Heroes []Hero `json:"Heroes" gorm:"many2many:hero_superpowers;association_foreignkey:ID;foreignkey:ID"`
}
