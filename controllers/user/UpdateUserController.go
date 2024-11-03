package controllers

import (
	"mensina-be/core/models"
	"mensina-be/core/useCases/user"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update user by ID
// @Tags User
// @Param id path int true "User ID"
// @Param user body models.User false "User object"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.User "Success"
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
	var _user models.User

	err = c.ShouldBindJSON(&_user)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "cannot bind JSON",
		})
		return
	}

	if int(_user.ID) != id {
		c.JSON(400, utils.ErrorResponse{
			Error: "request ID must be equal to the body object ID",
		})
		return
	}

	updatedUser, status, err := user.UpdateUser(&_user)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, updatedUser)
}
