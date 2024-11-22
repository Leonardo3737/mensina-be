package quizUseCase

import (
	"fmt"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/database"
	"mensina-be/database/models"
	"sync"
)

func AnswerCheck(answerId, questionId, userId int, quizRoutineChannel chan routines.RoutineCallback) (bool, *config.RestErr) {
	db := database.GetDatabase()

	var answer models.Answer

	if err := db.Preload("Question").First(&answer, answerId).Error; err != nil {
		return false, config.NewNotFoundErr(fmt.Sprintf("cannot find answer: %d", answerId))
	}

	var wg sync.WaitGroup

	var err *config.RestErr
	isCorrect := answer.IsCorrect && answer.QuestionId == uint(questionId)

	wg.Add(1)
	quizRoutineChannel <- func(quizSessions routines.QuizSessions) *sync.WaitGroup {
		if answer.QuestionId != uint(questionId) {
			err = config.NewBadRequestErr("this answer is not related to this question")
			return &wg
		}

		sessionKey := services.GetQuizSessionsKey(uint(userId), answer.Question.QuizId)
		quizSession, exist := quizSessions[sessionKey]

		fmt.Printf("quizID: %s\n", sessionKey)
		if !exist || quizSession.Total == 5 {
			err = config.NewNotFoundErr("this quiz is not in progress")
			return &wg
		}
		if quizSession.Questions[questionId] != dto.Unanswered {
			err = config.NewBadRequestErr("this question has already been answered")
			return &wg
		}

		quizSession.Total++

		if !isCorrect {
			quizSession.Questions[questionId] = dto.InCorrect
		} else {
			quizSession.Questions[questionId] = dto.Correct
			scoreToAdd := 10
			quizSession.Correct++

			if quizSession.Total > 3 && quizSession.Total == quizSession.Correct {
				scoreToAdd += quizSession.Total
			}

			quizSession.Score += scoreToAdd
		}

		if quizSession.Total == 5 {
			ApplyFinalBonus(quizSession)
		}
		return &wg
	}
	wg.Wait()
	return isCorrect, err
}

func ApplyFinalBonus(quizSession *dto.QuizSession) {
	accuracy := float64(quizSession.Correct) / float64(quizSession.Total)

	if accuracy == 1.0 {
		// 100% de acertos
		quizSession.Score = int(float64(quizSession.Score) * 1.2) // +20%
	} else if accuracy >= 0.8 {
		// 80% ou mais de acertos
		quizSession.Score = int(float64(quizSession.Score) * 1.1) // +10%
	}
}
