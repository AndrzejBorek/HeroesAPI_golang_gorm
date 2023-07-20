package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	Name        string       `json:"name" gorm:"Not Null;unique"`
	SuperPowers []SuperPower `json:"superPowers" gorm:"unique;many2many:hero_superpowers"`
	Villains    []*Villain   `json:"villains" gorm:"many2many:hero_villains"`
	Helpers     []Helper     `json:"helpers"`
	SuperTeamID uint         `json:"SuperTeamID"`
	SuperTeam   SuperTeam    `json:"-"`
}

func (Hero) TableName() string {
	return "Hero"
}
