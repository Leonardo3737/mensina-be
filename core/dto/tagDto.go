package dto

type CreateTagDto struct {
	Description string `json:"description" binding:"required"`
}
