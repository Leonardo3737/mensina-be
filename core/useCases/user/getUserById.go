package user

import (
	"mensina-be/core/models"
	"mensina-be/database"
)

func GetUserById(id int) (models.User, error) {
	db := database.GetDatabase()
	var user models.User

	err := db.First(&user, id).Error

	return user, err
}
