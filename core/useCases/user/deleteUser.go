package user

import (
	"mensina-be/core/models"
	"mensina-be/database"
)

func DeleteUser(id int) error {
	db := database.GetDatabase()

	err := db.Delete(&models.User{}, id).Error

	return err
}
