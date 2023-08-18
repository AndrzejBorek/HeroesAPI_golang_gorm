package requestModels

type CreateSuperPowerInput struct {
	Description string `json:"description" binding:"required"`
}
