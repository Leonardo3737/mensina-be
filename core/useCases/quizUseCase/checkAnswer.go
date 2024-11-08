package quizUseCase

import (
	"fmt"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/database"
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
	return isCorrect, err
}
