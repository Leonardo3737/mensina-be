package authUseCase

import (
	"fmt"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/core/services"
	"mensina-be/database"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Login(login *dto.InputLoginDto) (dto.OutputToken, int, error) {

	err := validate.Struct(login)
	if err != nil {
		return dto.OutputToken{}, 400, err
	}

	db := database.GetDatabase()

	var user models.User
	err = db.Where("user_name = ?", login.Username).First(&user).Error
	if err != nil {
		return dto.OutputToken{}, 404, fmt.Errorf("user not found")
	}

	if user.Password != services.SHA256Enconder(login.Password) {
		return dto.OutputToken{}, 401, fmt.Errorf("invalid password")
	}

	token, err := services.NewJWRService().GenerateToken(user.ID)
	if err != nil {
		return dto.OutputToken{}, 500, fmt.Errorf("internal server error")
	}

	return dto.OutputToken{Token: token}, 200, nil
}
