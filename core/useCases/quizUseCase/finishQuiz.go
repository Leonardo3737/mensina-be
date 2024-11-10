package quizUseCase

import (
	"fmt"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/core/useCases/rankUseCase"
	"mensina-be/database"
	"mensina-be/database/models"
	"sync"
)

func FinishQuiz(quizId, userId uint, quizRoutineChannel chan routines.RoutineCallback) {
	sessionQuizId := services.GetQuizSessionsKey(userId, quizId)

	quizRoutineChannel <- func(qs routines.QuizSessions) *sync.WaitGroup {
		quizSession, exist := qs[sessionQuizId]
		if !exist {
			fmt.Println("quiz não iniciado")
			return nil
		}
		defer delete(qs, sessionQuizId)

		if quizSession.Total == 5 {
			db := database.GetDatabase()
			userCompletedQuiz := models.UserCompletedQuiz{
				Score:  quizSession.Score,
				UserId: userId,
				QuizId: quizId,
			}
			err := db.Create(&userCompletedQuiz).Error
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			go rankUseCase.UpdateRank()
		}

		fmt.Println("quiz finalizado")
		return nil
	}
}
