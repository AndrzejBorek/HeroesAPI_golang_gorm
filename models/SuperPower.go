package models

import "gorm.io/gorm"

type SuperPower struct {
	gorm.Model
	ID          uint   `json:"ID" gorm:"primary_key"`
	Description string `json:"description" gorm:"Not Null"`
}
