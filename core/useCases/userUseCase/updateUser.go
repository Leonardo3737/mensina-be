package userUseCase

import (
	"errors"
	"fmt"
	"log"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/services"
	"mensina-be/database"

	"gorm.io/gorm"
)

func UpdateUser(user *dto.UpdateUserDto, id uint) (int, error) {
	db := database.GetDatabase()

	var existingUser models.User
	err := db.Where("user_name = ?", user.UserName).First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Erro inesperado (por exemplo, conexão com o banco)
		return 500, fmt.Errorf("cannot checking username")
	} else if err == nil {
		// Usuário já existe
		return 409, fmt.Errorf("username already exists")
	}

	fmt.Println(id)

	if user.Password != "" {
		log.Print("senha")
		user.Password = services.SHA256Enconder(user.Password)
	}

	res := db.
		Model(&existingUser).
		Where("id = ?", id).
		Updates(user)

	if res.Error != nil {
		return 500, fmt.Errorf("cannot update user")
	}

	if res.RowsAffected == 0 {
		return 404, fmt.Errorf("cannot find user. id: %d", id)
	}

	return 204, err
}
