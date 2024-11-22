package models

import "time"

type UserCompletedQuiz struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Score          int       `json:"score"`
	CorrectAnswers int       `json:"correctAnswers"`
	UserId         uint      `json:"userId"`
	User           User      `gorm:"foreignKey:UserId"`
	QuizId         uint      `json:"quizId"`
	Quiz           Quiz      `gorm:"foreignKey:QuizId"`
	CreatedAt      time.Time `json:"createdAt" gorm:"type:timestamp;default:current_timestamp"`
}
