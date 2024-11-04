package userController

import (
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Tags User
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Security BearerAuth
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "ID has to be integer",
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
