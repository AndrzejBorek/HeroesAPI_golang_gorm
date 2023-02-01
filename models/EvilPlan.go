package models

import "gorm.io/gorm"

type EvilPlan struct {
	gorm.Model
	ID          uint   `json:"ID" gorm:"primary_key"`
	Description string `json:"name" gorm:"Not Null"`
}
