package userUseCase

import (
	"mensina-be/database"
	"mensina-be/database/models"
)

func GetUserInfos(id uint) (models.User, error) {
	db := database.GetDatabase()
	var user models.User

	err := db.First(&user, id).Error

	return user, err
}
