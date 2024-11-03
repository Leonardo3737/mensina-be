package controllers

import (
	"mensina-be/core/models"
	"mensina-be/core/useCases/auth"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Authenticate
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.Login true "Authentication data"
// @Success 201 {object} models.LoginRes "Token"
// @Router /login [post]
func Login(c *gin.Context) {
	var _login models.Login

	err := c.ShouldBindJSON(&_login)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "cannot bind JSON",
		})
		return
	}

	res, status, err := auth.Login(&_login)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, res)
}
