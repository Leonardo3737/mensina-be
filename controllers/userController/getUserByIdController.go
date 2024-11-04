package userController

import (
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get user by ID
// @Tags User
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} models.User "Success"
// @Router /user/{id} [get]
func GetById(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "ID has to be integer",
		})
		return
	}
	user, err := userUseCase.GetUserById(id)

	if err != nil {
		c.JSON(404, utils.ErrorResponse{
			Error: "cannot find user",
		})
		return
	}

	c.JSON(200, user)
}
