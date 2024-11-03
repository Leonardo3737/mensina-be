package auth

import (
	"fmt"
	"mensina-be/core/models"
	"mensina-be/core/services"
	"mensina-be/database"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Login(login *models.Login) (models.LoginRes, int, error) {

	err := validate.Struct(login)
	if err != nil {
		return models.LoginRes{}, 400, err
	}

	db := database.GetDatabase()

	var user models.User
	err = db.Where("user_name = ?", login.Username).First(&user).Error
	if err != nil {
		return models.LoginRes{}, 404, fmt.Errorf("user not found")
	}

	if user.Password != services.SHA256Enconder(login.Password) {
		return models.LoginRes{}, 401, fmt.Errorf("invalid password")
	}

	token, err := services.NewJWRService().GenerateToken(user.ID)
	if err != nil {
		return models.LoginRes{}, 500, fmt.Errorf("internal server error")
	}

	return models.LoginRes{Token: token}, 200, nil
}
