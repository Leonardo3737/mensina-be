package controllers

import (
	"mensina-be/core/models"
	"mensina-be/core/useCases/user"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User "User"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var _user models.User

	err := c.ShouldBindJSON(&_user)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "cannot bind JSON",
		})
		return
	}

	newUser, status, err := user.CreateUser(&_user)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, newUser)
}
