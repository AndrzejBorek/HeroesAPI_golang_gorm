package requestModels

type CreateVillainInput struct {
	Name             string `json:"name" binding:"required"`
	VillainPowersIDs []uint `json:"superPowers" binding:"required"`
	HeroEnemiesIDs   []uint `json:"heroEnemies"`
	EvilPlan         []uint `json:"evilPlan"`
}
