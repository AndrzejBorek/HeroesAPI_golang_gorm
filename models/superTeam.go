package models

import "gorm.io/gorm"

type SuperTeam struct {
	gorm.Model
	Name   string `json:"Name" gorm:"UNIQUE"`
	Heroes []Hero `json:"HeroesIDs" gorm:"constraint:OnDelete:SET NULL;"`
}

func (SuperTeam) TableName() string {
	return "SuperTeam"
}
