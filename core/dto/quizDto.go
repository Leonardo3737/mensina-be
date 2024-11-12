package dto

type CreateQuizDto struct {
	Title string `json:"title" binding:"required"`
	TagID uint   `json:"tagId" binding:"required"`
}
