package tagUseCase

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"
)

func CreateTag(tag *dto.CreateTagDto) (models.Tag, *config.RestErr) {
	db := database.GetDatabase()

	// Criptografar senha
	newTag := models.Tag{
		Description: tag.Description,
	}

	err := db.Create(&newTag).Error

	if err != nil {
		return models.Tag{}, config.NewInternaErr("cannot create Tag")
	}

	return newTag, nil
}
