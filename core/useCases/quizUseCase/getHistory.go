package quizUseCase

import (
	"mensina-be/config"
	"mensina-be/database"
	"mensina-be/database/models"
)

func GetHistory(userId uint) ([]models.UserCompletedQuiz, *config.RestErr) {
	db := database.GetDatabase()

	var history []models.UserCompletedQuiz

	err := db.
		Preload("Quiz.Tag").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&history).Error

	if err != nil {
		return []models.UserCompletedQuiz{}, config.NewInternaErr("cannot find quizzes history")
	}

	return history, nil
}
