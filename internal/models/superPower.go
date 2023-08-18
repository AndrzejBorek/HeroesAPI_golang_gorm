package models

import "gorm.io/gorm"

type SuperPower struct {
	gorm.Model
	Description string `json:"description" gorm:"Not Null;unique"`
}

func (SuperPower) TableName() string {
	return "SuperPower"
}
