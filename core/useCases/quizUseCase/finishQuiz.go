package quizUseCase

import (
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/core/useCases/rankUseCase"
	"mensina-be/database"
	"mensina-be/database/models"
	"sync"
)

func FinishQuiz(quizId, userId uint, quizRoutineChannel chan routines.RoutineCallback) dto.QuizSession {
	sessionQuizId := services.GetQuizSessionsKey(userId, quizId)

	var wg sync.WaitGroup
	var _quizSession dto.QuizSession

	wg.Add(1)
	quizRoutineChannel <- func(qs routines.QuizSessions) *sync.WaitGroup {
		quizSession, exist := qs[sessionQuizId]
		if !exist {
			fmt.Println("quiz não iniciado")
			return &wg
		}
		_quizSession = *quizSession
		fmt.Printf("%+v", _quizSession)
		defer delete(qs, sessionQuizId)

		if quizSession.Total == 5 {
			db := database.GetDatabase()
			userCompletedQuiz := models.UserCompletedQuiz{
				CorrectAnswers: quizSession.Correct,
				Score:          quizSession.Score,
				UserId:         userId,
				QuizId:         quizId,
			}
			err := db.Create(&userCompletedQuiz).Error
			if err != nil {
				fmt.Println(err.Error())
				return &wg
			}
			go rankUseCase.UpdateRank()
		}

		fmt.Println("quiz finalizado")
		return &wg
	}
	wg.Wait()

	fmt.Printf("%+v", _quizSession)
	return _quizSession
}
