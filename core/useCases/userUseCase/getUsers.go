package userUseCase

import (
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"
)

func GetUsers() ([]dto.OutputUserDto, error) {
	db := database.GetDatabase()

	var users []dto.OutputUserDto

	err := db.
	Model(&models.User{}).
	Select("id, user_name, name").
	Scan(&users).Error

	return users, err
}
