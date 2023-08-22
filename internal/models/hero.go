package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	Name        string       `json:"name" gorm:"Not Null;unique"`
	SuperPowers []SuperPower `json:"superPowers" gorm:"many2many:hero_superpowers"`
	Villains    []*Villain   `json:"villains" gorm:"many2many:hero_villains"`
	SuperTeams  []*SuperTeam `json:"superTeams" gorm:"many2many:heroes_superteam"`
	Helpers     []*Helper    `json:"helpers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Hero) TableName() string {
	return "Hero"
}

type HeroNameID struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func SerializeHeroes(heroes []*Hero) []HeroNameID {
	serializedHeroes := make([]HeroNameID, len(heroes))
	for i, hero := range heroes {
		serializedHeroes[i] = HeroNameID{
			ID:   hero.ID,
			Name: hero.Name,
		}
	}
	return serializedHeroes
}
