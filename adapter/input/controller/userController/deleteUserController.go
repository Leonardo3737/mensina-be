package userController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Tags User
// @Success 204 "No Content"
// @Security BearerAuth
// @Router /user/ [delete]
func (controller userController) DeleteUser(c *gin.Context) {
	id, err := utils.GetUserIdByToken(c)

	if err != nil {
		restErr := config.NewUnauthorizedErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	if restErr := userUseCase.DeleteUser(id); restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(204)
}
