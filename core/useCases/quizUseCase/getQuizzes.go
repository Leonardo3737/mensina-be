package quizUseCase

import (
	"log"
	"mensina-be/core/routines"
	"mensina-be/core/services"
	"mensina-be/database"
	"mensina-be/database/models"
	"sync"
)

func GetQuizzes(tagId string, inProgress bool, quizRoutineChannel chan routines.RoutineCallback) ([]models.Quiz, error) {
	db := database.GetDatabase()
	var wg sync.WaitGroup

	var quizzes []models.Quiz
	var quizzesIdInProgres []uint

	wg.Add(1)
	quizRoutineChannel <- func(qs routines.QuizSessions) *sync.WaitGroup {
		for sessiokey, _ := range qs {
			quizId, err := services.ExtractQuizId(sessiokey)
			if err != nil {
				log.Fatal(err)
			}
			quizzesIdInProgres = append(quizzesIdInProgres, quizId)
		}
		return &wg
	}
	wg.Wait()

	query := db.Preload("Tag")

	if tagId != "" {
		query = query.Where("tag_id = ?", tagId)
	}

	if inProgress {
		if len(quizzesIdInProgres) > 0 {
			query = query.Where("id = ?", quizzesIdInProgres[0])

			for _, quizId := range quizzesIdInProgres[1:] {
				query = query.Or("id = ?", quizId)
			}
		} else {
			query = query.Where("false")
		}
	} else {
		for _, quizId := range quizzesIdInProgres {
			query = query.Where("id != ?", quizId)
		}
	}

	err := query.Find(&quizzes).Error

	return quizzes, err
}
