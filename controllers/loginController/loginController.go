package loginController

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/useCases/loginUseCase"
	"net/http"

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
		restErr := config.NewBadRequestErr("cannot bind JSON")
		c.JSON(restErr.Code, restErr)
		return
	}

	res, restErr := loginUseCase.Login(&_login)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, res)
}
