package userUseCase

import (
	"errors"
	"fmt"
	"log"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/services"
	"mensina-be/database"
	"mensina-be/database/models"

	"gorm.io/gorm"
)

func UpdateUser(user *dto.UpdateUserDto, id uint) *config.RestErr {
	db := database.GetDatabase()

	if user.Password != "" && len(user.Password) < 6 {
		return config.NewBadRequestErr("password must be longer than 6 characters")
	}

	if user.Name != "" && len(user.Name) < 2 {
		return config.NewBadRequestErr("name must be longer than 2 characters")
	}

	if user.UserName != "" && len(user.UserName) < 3 {
		return config.NewBadRequestErr("username must be longer than 3 characters")
	}

	var existingUser models.User
	err := db.Where("user_name = ?", user.UserName).First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Erro inesperado (por exemplo, conexão com o banco)
		return config.NewInternaErr("cannot checking username")
	} else if err == nil && existingUser.ID != id {
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

	return nil
}
