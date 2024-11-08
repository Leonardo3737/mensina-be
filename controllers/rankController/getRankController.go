package rankController

import (
	"mensina-be/core/useCases/rankUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get rank
// @Tags Rank
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.RankDto "Success"
// @Router /rank [get]
func GetRank(c *gin.Context) {
	rank, status, err := rankUseCase.GetRank()
	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: "cannot list users",
		})
		return
	}

	c.JSON(status, rank)
}
