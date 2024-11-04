package userUseCase

import (
	"mensina-be/core/models"
	"mensina-be/database"
)

func GetUsers() ([]models.User, error) {
	db := database.GetDatabase()

	var users []models.User

	err := db.Find(&users).Error

	return users, err
}
