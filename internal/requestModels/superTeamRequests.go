package requestModels

type CreateSuperTeamInput struct {
	Name      string `json:"Name" binding:"required"`
	HeroesIDs []uint `json:"HeroesIDs"  binding:"required"`
}

type AddHeroToSuperTeamInput struct {
	HeroesIDs []uint `json:"heroesIDs" binding:"required"`
}
