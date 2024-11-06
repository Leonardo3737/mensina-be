package quizUseCase

import (
	"errors"
	"fmt"
	"mensina-be/core/models"
	"mensina-be/database"

	"gorm.io/gorm"
)

type ChRealizeQuiz struct {
	Scored  bool
	UserId  uint
	QuizzId uint
}

type ResRealizeQuiz struct {
	QuizzId uint
	Hits    int
	Total   int
}

var Sections = make(map[int]*ResRealizeQuiz)

func StartQuiz(userId, quizId uint) (int, error) {
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

	return 200, nil
}

func updateQuizState(ch chan ChRealizeQuiz) {

	for {
		score := <-ch
		section, exist := Sections[int(score.UserId)]

		if !exist {
			fmt.Println("Iniciando Quiz")
			section.Total = 0
			section.Hits = 0
			section.QuizzId = score.QuizzId
			break
		}

		section.Total++

		if score.Scored {
			section.Hits++
		}

		fmt.Printf("Respondidos: %x\n", section.Total)
		fmt.Printf("%x\n", section.Total)
		if section.Total == 5 {
			break
		}
	}
	fmt.Printf("Finalizando Quiz")
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
