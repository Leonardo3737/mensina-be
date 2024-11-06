package userController

import (
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
		c.JSON(401, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := userUseCase.GetUserInfos(id)

	if err != nil {
		c.JSON(404, utils.ErrorResponse{
			Error: "cannot find user",
		})
		return
	}

	c.JSON(200, user)
}
