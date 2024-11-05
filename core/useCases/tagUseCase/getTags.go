package tagUseCase

import (
	"mensina-be/core/models"
	"mensina-be/database"
)

func GetTags() ([]models.Tag, error) {
	db := database.GetDatabase()

	var tags []models.Tag

	err := db.Find(&tags).Error

	return tags, err
}
