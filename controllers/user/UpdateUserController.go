package controllers

import (
	"mensina-be/core/dto"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update user by ID
// @Tags User
// @Param id path int true "User ID"
// @Param user body dto.UpdateUserDto true "User object"
// @Security BearerAuth
// @Success 204 "Success"
// @Router /user/{id} [put]
func UpdateUser(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "ID has to be integer",
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
