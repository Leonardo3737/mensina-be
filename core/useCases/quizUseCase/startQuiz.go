package quizUseCase

import (
	"errors"
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/database"

	"gorm.io/gorm"
)

func StartQuiz(userId, quizId uint, quizRoutineChannel chan dto.QuizRoutineChannel) (int, error) {
	db := database.GetDatabase()

	var user models.User

	if status, err := getEntity(&user, userId, db); err != nil {
		return status, err
	}

	fmt.Println(user.Name)

	var quiz models.Quiz
	if status, err := getEntity(&quiz, quizId, db); err != nil {
		return status, err
	}
	fmt.Println(quiz.Title)

	select {
	case sectionCh := <-quizRoutineChannel:
		fmt.Println(sectionCh)
		return 201, nil
	default:
		quizRoutineChannel <- dto.QuizRoutineChannel{
			Score:   0,
			UserId:  userId,
			QuizzId: quizId,
		}

		return 200, nil
	}
}

func getEntity(model interface{}, id uint, db *gorm.DB) (int, error) {

	if err := db.First(model, id).Error; err != nil {
		// Verifica se o erro indica que o registro não foi encontrado
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, errors.New("quiz not found")
		}
		// Caso contrário, retorna o erro do banco de dados
		return 500, errors.New("internal server error")
	}
	return 200, nil
}
