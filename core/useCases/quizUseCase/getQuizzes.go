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
	var quizzesIdInProgress []uint

	// Coletar IDs de quizzes em progresso
	wg.Add(1)
	quizRoutineChannel <- func(qs routines.QuizSessions) *sync.WaitGroup {
		for sessionKey := range qs {
			quizId, err := services.ExtractQuizId(sessionKey)
			if err != nil {
				log.Fatal(err)
			}
			quizzesIdInProgress = append(quizzesIdInProgress, quizId)
		}
		return &wg
	}
	wg.Wait()

	// Inicializar consulta
	query := db.Preload("Tag")

	// Filtrar por tagId
	if tagId != "" {
		query = query.Where("tag_id = ?", tagId)
	}

	// Filtrar por quizzes em progresso
	if inProgress {
		if len(quizzesIdInProgress) > 0 {
			// Agrupar condição OR
			query = query.Where("id IN (?)", quizzesIdInProgress)
		} else {
			// Nenhum quiz em progresso
			query = query.Where("false")
		}
	} else {
		// Excluir quizzes em progresso
		if len(quizzesIdInProgress) > 0 {
			query = query.Where("id NOT IN (?)", quizzesIdInProgress)
		}
	}

	// Executar consulta
	err := query.Find(&quizzes).Error

	return quizzes, err
}
