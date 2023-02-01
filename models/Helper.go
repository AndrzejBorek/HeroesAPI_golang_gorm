package models

import "gorm.io/gorm"

type Helper struct {
	gorm.Model
	ID     uint   `json:"ID" gorm:"primary_key"`
	Name   string `json:"name" gorm:"unique;not null;"`
	HeroID int
	Hero   Hero `json:"hero" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
