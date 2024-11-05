package quizUseCase

import (
	"fmt"
	"mensina-be/core/models"
	"mensina-be/database"
)

func AnswerCheck(answerId, questionId int) (bool, error) {
	db := database.GetDatabase()

	var answer models.Answer

	if err := db.First(&answer, answerId).Error; err != nil {
		return false, fmt.Errorf("cannot find answer: %d", answerId)
	}

	return answer.IsCorrect && answer.QuestionId == uint(questionId), nil
}
