package quizUseCase

import (
	"mensina-be/core/models"
	"mensina-be/database"
)

func GetQuizzes() ([]models.Quiz, error) {
	db := database.GetDatabase()

	var quizzes []models.Quiz

	err := db.Preload("Tag").Find(&quizzes).Error

	return quizzes, err
}
