package quizUseCase

import (
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"
)

func GetQuestionByQuiz(id int) ([]dto.OutputQuestionDto, error) {
	db := database.GetDatabase()
	var questions []models.Question

	err := db.Preload("Answers").Where("quiz_id = ?", id).Find(&questions).Error

	if err != nil {
		return []dto.OutputQuestionDto{}, fmt.Errorf("cannot find questions for quiz ID %d", id)
	}

	questionsDto := make([]dto.OutputQuestionDto, 0, len(questions))

	for _, question := range questions {
		answersDto := make([]dto.OutputAnswerDto, 0, len(question.Answers))

		for _, answer := range question.Answers {
			answersDto = append(answersDto, dto.OutputAnswerDto{
				ID:          answer.ID,
				Description: answer.Description,
			})
		}

		questionsDto = append(questionsDto, dto.OutputQuestionDto{
			ID:          question.ID,
			Title:       question.Title,
			Description: question.Description,
			Answers:     answersDto,
		})
	}

	return questionsDto, err
}
