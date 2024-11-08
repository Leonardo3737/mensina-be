package userUseCase

import (
	"mensina-be/config"
	"mensina-be/core/models"
	"mensina-be/database"
)

func DeleteUser(id uint) *config.RestErr {
	db := database.GetDatabase()

	result := db.Delete(&models.User{}, id)

	if result.Error != nil {
		return config.NewNotFoundErr("cannot delete user")
	}

	if int(result.RowsAffected) == 0 {
		return config.NewNotFoundErr("user not found")
	}

	return nil
}
