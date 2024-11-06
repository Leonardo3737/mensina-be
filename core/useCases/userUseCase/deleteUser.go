package userUseCase

import (
	"fmt"
	"mensina-be/core/models"
	"mensina-be/database"
)

func DeleteUser(id uint) error {
	db := database.GetDatabase()

	result := db.Delete(&models.User{}, id)

	if int(result.RowsAffected) == 0 {
		return fmt.Errorf("error: user not found")
	}

	return result.Error
}
