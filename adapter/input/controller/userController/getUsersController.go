package userController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Tags User
// @Produce json
// @Success 200 {array} models.User "Success"
// @Router /user [get]
func (controller userController) GetUsers(c *gin.Context) {
	users, err := userUseCase.GetUsers()

	if err != nil {
		restErr := config.NewInternaErr("cannot list users")
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(200, users)
}
