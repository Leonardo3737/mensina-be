package userController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get user kpi
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserKpiDto "Success"
// @Router /user/kpi [get]
func GetUserKpi(c *gin.Context) {
	userId, err := utils.GetUserIdByToken(c)

	if err != nil {
		restErr := config.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	userKpi, restErr := userUseCase.GetUserKpi(userId)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, userKpi)
}
