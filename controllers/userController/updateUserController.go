package userController

import (
	"mensina-be/config"
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
		restErr := config.NewUnauthorizedErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	var _user dto.UpdateUserDto

	if err := c.ShouldBindJSON(&_user); err != nil {
		restErr := config.NewBadRequestErr("some field are incorrect")
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := userUseCase.UpdateUser(&_user, id)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	c.Status(204)
}
