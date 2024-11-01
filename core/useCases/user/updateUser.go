package user

import (
	"errors"
	"fmt"
	"mensina-be/core/models"
	"mensina-be/database"

	"gorm.io/gorm"
)

func UpdateUser(user *models.User) (models.User, int, error) {

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

	err = db.Save(user).Error

	if err != nil {
		return models.User{}, 500, fmt.Errorf("cannot create user")
	}

	return *user, 204, err
}
