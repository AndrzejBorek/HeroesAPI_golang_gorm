package models

import "gorm.io/gorm"

type SuperTeam struct {
	gorm.Model
	//Heroes []Hero `json:"Heroes" gorm:"many2many:hero_superpowers;association_foreignkey:ID;foreignkey:ID"`
}

func (SuperTeam) TableName() string {
	return "SuperTeam"
}
