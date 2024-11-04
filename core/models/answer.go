package models

type Answer struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Description string   `json:"description"`
	IsCorrect   bool     `json:"isCorrect"`
	QuestionId  uint     `json:"questionId"`
	Question    Question `gorm:"foreignKey:QuestionId"`
}
