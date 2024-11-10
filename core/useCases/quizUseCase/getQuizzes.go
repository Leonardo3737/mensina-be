package quizUseCase

import (
	"mensina-be/database"
	"mensina-be/database/models"
)

func GetQuizzes(tagId string) ([]models.Quiz, error) {
	db := database.GetDatabase()

	var quizzes []models.Quiz

	var err error

	if tagId == "" {
		err = db.Preload("Tag").Find(&quizzes).Error
	} else {
		err = db.Preload("Tag").Where("tag_id = ?", tagId).Find(&quizzes).Error
	}

	return quizzes, err
}
