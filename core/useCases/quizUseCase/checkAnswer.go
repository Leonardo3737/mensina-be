package quizUseCase

import (
	"errors"
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/database"
	"sync"
)

func AnswerCheck(answerId, questionId, userId int, quizRoutineChannel chan routines.RoutineCallback) (bool, int, error) {
	db := database.GetDatabase()

	var answer models.Answer

	if err := db.Preload("Question").First(&answer, answerId).Error; err != nil {
		return false, 404, fmt.Errorf("cannot find answer: %d", answerId)
	}

	var wg sync.WaitGroup

	status := 200
	var err error = nil
	isCorrect := answer.IsCorrect && answer.QuestionId == uint(questionId)

	wg.Add(1)
	quizRoutineChannel <- func(quizSessions routines.QuizSessions) *sync.WaitGroup {
		if answer.QuestionId != uint(questionId) {
			err = errors.New("this answer is not related to this question")
			status = 400
			return &wg
		}

		sessionKey := services.GetQuizSessionsKey(uint(userId), answer.Question.QuizId)
		quizSession, exist := quizSessions[sessionKey]

		fmt.Printf("quizID: %s\n", sessionKey)
		if !exist || quizSession.Total == 5 {
			err = errors.New("this quiz is not in progress")
			status = 404
			return &wg
		}
		if quizSession.Questions[questionId] != dto.Unanswered {
			err = errors.New("this question has already been answered")
			status = 400
			return &wg
		}

		quizSession.Total++
		if quizSession.Total == 5 {
			defer func() {
				go FinishQuiz(quizSession.QuizzId, quizSession.UserId, quizRoutineChannel)
			}()
		}

		if !isCorrect {
			quizSession.Questions[questionId] = dto.InCorrect
			return &wg
		}
		quizSession.Questions[questionId] = dto.Correct

		scoreToAdd := 2
		quizSession.Correct++

		if quizSession.Total > 3 && quizSession.Total == quizSession.Correct {
			scoreToAdd = 3
		}
		quizSession.Score += scoreToAdd

		return &wg
	}
	wg.Wait()
	return isCorrect, status, err
}
