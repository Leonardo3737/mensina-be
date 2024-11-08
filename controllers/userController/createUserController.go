package userController

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/useCases/userUseCase"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User data"
// @Success 201 {object} models.User "User"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var _user dto.CreateUserDto

	if err := c.ShouldBindJSON(&_user); err != nil {
		restErr := config.NewBadRequestErr("some field are incorrect")
		c.JSON(restErr.Code, restErr)
		return
	}

	newUser, restErr := userUseCase.CreateUser(&_user)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(201, newUser)
}
