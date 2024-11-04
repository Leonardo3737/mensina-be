package controllers

import (
	"mensina-be/core/dto"
	"mensina-be/core/useCases/authUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Authenticate
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.InputLoginDto true "Authentication data"
// @Success 201 {object} dto.OutputToken "Token"
// @Router /login [post]
func Login(c *gin.Context) {
	var _login dto.InputLoginDto

	err := c.ShouldBindJSON(&_login)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "cannot bind JSON",
		})
		return
	}

	res, status, err := authUseCase.Login(&_login)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, res)
}
