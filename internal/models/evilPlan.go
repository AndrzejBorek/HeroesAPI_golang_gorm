package models

import "gorm.io/gorm"

type EvilPlan struct {
	gorm.Model
	Description string   `json:"name" gorm:"Not Null"`
	VillainId   uint     `json:"villainId"`
	Villain     *Villain `json:"villain" gorm:"Not Null"`
}

func (EvilPlan) TableName() string {
	return "EvilPlan"
}
