package quizUseCase

import (
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/database"
	"sync"
)

func AnswerCheck(answerId, questionId, userId int, quizRoutineChannel chan routines.RoutineCallback) (bool, error) {
	db := database.GetDatabase()

	var answer models.Answer

	if err := db.Preload("Question").First(&answer, answerId).Error; err != nil {
		return false, fmt.Errorf("cannot find answer: %d", answerId)
	}

	isCorrect := answer.IsCorrect && answer.QuestionId == uint(questionId)

	quizRoutineChannel <- func(quizSessions routines.QuizSessions) *sync.WaitGroup {
		if answer.QuestionId != uint(questionId) {
			fmt.Println("essa resposta não é referente a essa pergunta")
			return nil
		}

		sessionKey := services.GetQuizSessionsKey(uint(userId), answer.Question.QuizId)
		quizSession, exist := quizSessions[sessionKey]

		fmt.Printf("quizID: %s\n", sessionKey)
		if !exist || quizSession.Total == 5 || quizSession.Questions[questionId] != dto.Unanswered {
			fmt.Printf("Quiz não iniciado ou finalizado, ou questão respondida, quizID: %s\n", sessionKey)
			if exist {
				fmt.Printf("status da resposta: %s\n", quizSession.Questions[questionId])
			}
			return nil
		}

		quizSession.Total++

		if !isCorrect {
			quizSession.Questions[questionId] = dto.InCorrect
			fmt.Println("Resposta incorreta")
			return nil
		}

		quizSession.Questions[questionId] = dto.Correct
		fmt.Println("Resposta correta")

		quizSession.Score++

		return nil
	}

	return isCorrect, nil
}
