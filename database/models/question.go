package models

type Question struct {
	ID      uint     `json:"id" gorm:"primaryKey"`
	Title   string   `json:"title"`
	QuizId  uint     `json:"quizId"`
	Quiz    Quiz     `gorm:"foreignKey:QuizId"`
	Answers []Answer `json:"quizzes" gorm:"foreignKey:QuestionId"`
}
