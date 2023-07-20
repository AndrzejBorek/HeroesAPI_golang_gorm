package models

import "gorm.io/gorm"

type Helper struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique;not null"`
	HeroID uint   `json:"heroID"`
	Hero   Hero   `json:"-"`
}

func (Helper) TableName() string {
	return "Helper"
}
