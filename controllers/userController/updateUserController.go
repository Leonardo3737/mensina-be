package userController

import (
	"mensina-be/core/dto"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Update user by ID
// @Tags User
// @Param user body dto.UpdateUserDto true "User object"
// @Security BearerAuth
// @Success 204 "Success"
// @Router /user/ [put]
func UpdateUser(c *gin.Context) {
	id, err := utils.GetUserIdByToken(c)

	if err != nil {
		c.JSON(401, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var _user dto.UpdateUserDto

	err = c.ShouldBindJSON(&_user)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "cannot bind JSON",
		})
		return
	}

	status, err := userUseCase.UpdateUser(&_user, id)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.Status(status)
}
