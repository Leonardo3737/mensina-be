package userUseCase

import (
	"errors"
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/services"
	"mensina-be/database"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func CreateUser(user *dto.CreateUserDto) (models.User, int, error) {

	err := validate.Struct(user)
	if err != nil {
		return models.User{}, 400, err
	}

	db := database.GetDatabase()

	var existingUser models.User
	err = db.Where("user_name = ?", user.UserName).First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Erro inesperado (por exemplo, conexão com o banco)
		return models.User{}, 500, fmt.Errorf("cannot checking username")
	} else if err == nil {
		// Usuário já existe
		return models.User{}, 409, fmt.Errorf("username already exists")
	}

	// Criptografar senha
	newUser := models.User{
		UserName: user.UserName,
		Name:     user.Name,
		Password: services.SHA256Enconder(user.Password),
	}

	err = db.Create(newUser).Error

	if err != nil {
		return models.User{}, 500, fmt.Errorf("cannot create user")
	}

	return newUser, 201, err
}
