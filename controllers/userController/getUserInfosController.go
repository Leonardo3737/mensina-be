package userController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get user info with Token
// @Tags User
// @Produce json
// @Success 200 {object} models.User "Success"
// @Security BearerAuth
// @Router /user/user_infos [get]
func GetUserInfos(c *gin.Context) {
	id, err := utils.GetUserIdByToken(c)

	if err != nil {
		restErr := config.NewUnauthorizedErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	user, err := userUseCase.GetUserInfos(id)

	if err != nil {
		restErr := config.NewNotFoundErr("cannot find user")
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(200, user)
}
