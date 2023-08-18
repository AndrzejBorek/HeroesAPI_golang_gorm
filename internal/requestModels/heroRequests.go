package requestModels

type CreateHeroInput struct {
	Name           string `json:"name" binding:"required"`
	SuperPowersIDs []uint `json:"superPowers" binding:"required"`
	VillainsIDs    []uint `json:"villains"`
	HelpersIDs     []uint `json:"helpers"`
	SuperTeamsIDs  []uint `json:"superTeams"`
}
