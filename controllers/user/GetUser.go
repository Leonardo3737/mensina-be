package controllers

import (
	"mensina-be/core/useCases/user"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := user.GetUsers()

	if err != nil {
		c.JSON(500, utils.ErrorResponse{
			Error: "cannot list users",
		})
		return
	}

	c.JSON(200, users)
}
