package quizUseCase

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"
)

func CreateQuiz(quiz *dto.CreateQuizDto) (models.Quiz, *config.RestErr) {
	db := database.GetDatabase()

	// Criptografar senha
	newQuiz := models.Quiz{
		Title: quiz.Title,
		TagID: quiz.TagID,
	}

	err := db.Create(&newQuiz).Error

	if err != nil {
		return models.Quiz{}, config.NewInternaErr("cannot create quiz")
	}

	return newQuiz, nil
}
