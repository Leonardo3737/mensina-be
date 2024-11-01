package controllers

import (
	"mensina-be/core/useCases/user"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "ID has to be integer",
		})
		return
	}
	err = user.DeleteUser(id)

	if err != nil {
		c.JSON(404, utils.ErrorResponse{
			Error: "cannot find user",
		})
		return
	}

	c.Status(204)
}
