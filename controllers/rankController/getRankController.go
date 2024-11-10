package rankController

import (
	"mensina-be/core/useCases/rankUseCase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get rank
// @Tags Rank
// @Produce json
// @Param update_rank query boolean false "Update Rank"
// @Security BearerAuth
// @Success 200 {array} models.Rank "Success"
// @Router /rank [get]
func GetRank(c *gin.Context) {
	updateRank := c.Query("update_rank") == "true"

	rank, restErr := rankUseCase.GetRank(updateRank)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, rank)
}
