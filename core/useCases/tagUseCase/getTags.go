package tagUseCase

import (
	"mensina-be/database"
	"mensina-be/database/models"
)

func GetTags() ([]models.Tag, error) {
	db := database.GetDatabase()

	var tags []models.Tag

	err := db.Find(&tags).Error

	return tags, err
}
