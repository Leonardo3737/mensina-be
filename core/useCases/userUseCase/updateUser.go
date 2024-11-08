package userUseCase

import (
	"errors"
	"fmt"
	"log"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/services"
	"mensina-be/database"

	"gorm.io/gorm"
)

func UpdateUser(user *dto.UpdateUserDto, id uint) *config.RestErr {
	db := database.GetDatabase()

	var existingUser models.User
	err := db.Where("user_name = ?", user.UserName).First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Erro inesperado (por exemplo, conexão com o banco)
		return config.NewInternaErr("cannot checking username")
	} else if err == nil {
		// Usuário já existe
		return config.NewConflictErr("username already exists")
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
		return config.NewInternaErr("cannot update user")
	}

	if res.RowsAffected == 0 {
		return config.NewNotFoundErr(fmt.Sprintf("cannot find user. id: %d", id))
	}

	return nil
}
