package models

type UserCompletedQuiz struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	Score  int  `json:"score"`
	UserId uint `json:"userId"`
	User   Quiz `gorm:"foreignKey:UserId"`
	QuizId uint `json:"quizId"`
	Quiz   Quiz `gorm:"foreignKey:QuizId"`
}
