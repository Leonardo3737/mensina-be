package loginUseCase

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/services"
	"mensina-be/database"
	"mensina-be/database/models"
)

func Login(login *dto.InputLoginDto) (dto.OutputToken, *config.RestErr) {
	db := database.GetDatabase()

	var user models.User
	err := db.Where("user_name = ?", login.Username).First(&user).Error
	if err != nil {
		return dto.OutputToken{}, config.NewNotFoundErr("user not found")
	}

	if user.Password != services.SHA256Enconder(login.Password) {
		return dto.OutputToken{}, config.NewUnauthorizedErr("invalid password")
	}

	token, err := services.NewJWRService().GenerateToken(user.ID)
	if err != nil {
		return dto.OutputToken{}, config.NewInternaErr("error generating token")
	}

	return dto.OutputToken{Token: token}, nil
}
