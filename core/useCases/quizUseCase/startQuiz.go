package quizUseCase

import (
	"errors"
	"fmt"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/database"
	"sync"

	"gorm.io/gorm"
)

func StartQuiz(userId, quizId uint, quizRoutineChannel chan routines.RoutineCallback) (dto.QuizSession, *config.RestErr) {
	db := database.GetDatabase()

	var user models.User
	if err := getEntity(&user, userId, db); err != nil {
		return dto.QuizSession{}, err
	}

	var quiz models.Quiz
	if err := getEntity(&quiz, quizId, db); err != nil {
		return dto.QuizSession{}, err
	}

	var wg sync.WaitGroup
	var _quizSession dto.QuizSession

	wg.Add(1)
	quizRoutineChannel <- func(sessions routines.QuizSessions) *sync.WaitGroup {
		sessionKey := services.GetQuizSessionsKey(userId, quizId)

		quizSession, exist := sessions[sessionKey]

		if !exist {
			fmt.Printf("Iniciando Quiz, userId: %d | quizId: %d\n", userId, quizId)

			quizSession = &dto.QuizSession{
				Total:     0,
				Score:     0,
				UserId:    userId,
				QuizzId:   quizId,
				Questions: make(map[int]dto.Status),
			}
			sessions[sessionKey] = quizSession
		}
		_quizSession = *quizSession
		return &wg
	}

	wg.Wait()

	return _quizSession, nil
}

func getEntity(model interface{}, id uint, db *gorm.DB) *config.RestErr {

	if err := db.First(model, id).Error; err != nil {
		// Verifica se o erro indica que o registro não foi encontrado
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return config.NewNotFoundErr("quiz not found")
		}
		// Caso contrário, retorna o erro do banco de dados
		return config.NewInternaErr("internal server error")
	}
	return nil
}
