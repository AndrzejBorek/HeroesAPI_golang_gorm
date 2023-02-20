package models

import "gorm.io/gorm"

type EvilPlan struct {
	gorm.Model
	Description string `json:"name" gorm:"Not Null"`
}

func (EvilPlan) TableName() string {
	return "EvilPlan"
}
