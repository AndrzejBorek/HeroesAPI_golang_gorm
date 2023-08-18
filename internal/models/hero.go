package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	Name        string       `json:"name" gorm:"Not Null;unique"`
	SuperPowers []SuperPower `json:"superPowers" gorm:"many2many:hero_superpowers"`
	Villains    []*Villain   `json:"villains" gorm:"many2many:hero_villains"`
	SuperTeams  []*SuperTeam `json:"superTeams" gorm:"many2many:heroes_superteam"`
	Helpers     []*Helper    `json:"helpers"`
}

func (Hero) TableName() string {
	return "Hero"
}
