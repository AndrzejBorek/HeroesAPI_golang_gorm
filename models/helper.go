package models

import "gorm.io/gorm"

type Helper struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique;not null"`
	HeroID uint   `json:"heroID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Hero   Hero   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Helper) TableName() string {
	return "Helper"
}
