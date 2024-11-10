package userUseCase

import (
	"errors"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/services"
	"mensina-be/database"
	"mensina-be/database/models"

	"gorm.io/gorm"
)

func CreateUser(user *dto.CreateUserDto) (models.User, *config.RestErr) {
	db := database.GetDatabase()

	var existingUser models.User
	err := db.Where("user_name = ?", user.UserName).First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Erro inesperado (por exemplo, conexão com o banco)
		return models.User{}, config.NewInternaErr("cannot checking username")
	} else if err == nil {
		// Usuário já existe
		return models.User{}, config.NewConflictErr("username already exists")
	}

	// Criptografar senha
	newUser := models.User{
		UserName: user.UserName,
		Name:     user.Name,
		Password: services.SHA256Enconder(user.Password),
	}

	err = db.Create(&newUser).Error

	if err != nil {
		return models.User{}, config.NewInternaErr("cannot create user")
	}

	return newUser, nil
}
