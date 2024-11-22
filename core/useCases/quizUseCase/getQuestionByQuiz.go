package quizUseCase

import (
	"fmt"
	"math/rand"
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"
	"time"
)

func GetQuestionByQuiz(id int) ([]dto.OutputQuestionDto, error) {
	db := database.GetDatabase()
	var questions []models.Question

	err := db.Preload("Answers").Where("quiz_id = ?", id).Find(&questions).Error

	if err != nil {
		return []dto.OutputQuestionDto{}, fmt.Errorf("cannot find questions for quiz ID %d", id)
	}

	questionsDto := make([]dto.OutputQuestionDto, 0, len(questions))

	rand.Seed(time.Now().UnixNano())

	for _, question := range questions {
		answersDto := make([]dto.OutputAnswerDto, 0, len(question.Answers))

		for _, answer := range question.Answers {
			answersDto = append(answersDto, dto.OutputAnswerDto{
				ID:          answer.ID,
				Description: answer.Description,
			})
		}

		// Embaralha as respostas
		rand.Shuffle(len(answersDto), func(i, j int) {
			answersDto[i], answersDto[j] = answersDto[j], answersDto[i]
		})

		questionsDto = append(questionsDto, dto.OutputQuestionDto{
			ID:      question.ID,
			Title:   question.Title,
			Answers: answersDto,
		})

	}

	// Embaralha as perguntas
	rand.Shuffle(len(questionsDto), func(i, j int) {
		questionsDto[i], questionsDto[j] = questionsDto[j], questionsDto[i]
	})

	return questionsDto, err
}
