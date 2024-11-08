package rankController

import (
	"mensina-be/core/useCases/rankUseCase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get rank
// @Tags Rank
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.RankDto "Success"
// @Router /rank [get]
func GetRank(c *gin.Context) {
	rank, restErr := rankUseCase.GetRank()
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, rank)
}
