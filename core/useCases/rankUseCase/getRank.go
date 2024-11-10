package rankUseCase

import (
	"mensina-be/config"
	"mensina-be/core/models"
	"mensina-be/database"
)

func GetRank(updateRank bool) ([]models.Rank, *config.RestErr) {
	db := database.GetDatabase()

	var rank []models.Rank

	if updateRank {
		UpdateRank()
	}

	err := db.Preload("User").Preload("BestScoreQuiz.Tag").Find(&rank).Error

	if err != nil {
		return rank, config.NewInternaErr("cannot list rank")
	}
	return rank, nil
}
