package rankUseCase

import (
	"mensina-be/config"
	"mensina-be/core/services"
	"mensina-be/database"
	"mensina-be/database/models"

	"gorm.io/gorm"
)

func GetRank(updateRank bool, page, perPage int) ([]models.Rank, *config.RestErr) {
	db := database.GetDatabase()

	var rank []models.Rank

	if updateRank {
		UpdateRank()
	}

	if perPage != 0 {
		db = services.Paginate(page, perPage, db)
	}

	err := db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "user_name")
		}).
		Preload("BestScoreQuiz.Tag").
		Order("total_score DESC").
		Find(&rank).
		Error

	if err != nil {
		return rank, config.NewInternaErr("cannot list rank")
	}
	return rank, nil
}
