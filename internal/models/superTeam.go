package models

import "gorm.io/gorm"

type SuperTeam struct {
	gorm.Model
	Name   string  `json:"name" gorm:"Not Null;unique"`
	Heroes []*Hero `json:"heroes"  gorm:"many2many:heroes_superteam"`
}

func (SuperTeam) TableName() string {
	return "SuperTeam"
}
