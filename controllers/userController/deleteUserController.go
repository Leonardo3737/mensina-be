package userController

import (
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Tags User
// @Success 204 "No Content"
// @Security BearerAuth
// @Router /user/ [delete]
func DeleteUser(c *gin.Context) {
	id, err := utils.GetUserIdByToken(c)

	if err != nil {
		c.JSON(401, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = userUseCase.DeleteUser(id)

	if err != nil {
		c.JSON(404, utils.ErrorResponse{
			Error: "cannot find user",
		})
		return
	}

	c.Status(204)
}
