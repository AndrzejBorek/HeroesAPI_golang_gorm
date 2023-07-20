package models

import "gorm.io/gorm"

type EvilPlan struct {
	gorm.Model
	Description string `json:"name" gorm:"Not Null"`
	VillainID   uint   `json:"villains"`
}

func (EvilPlan) TableName() string {
	return "EvilPlan"
}
