package rankController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/rankUseCase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get rank
// @Tags Rank
// @Produce json
// @Param update_rank query boolean false "Update Rank"
// @Param page query int false "Page"
// @Param perPage query int false "Per page"
// @Security BearerAuth
// @Success 200 {array} models.Rank "Success"
// @Router /rank [get]
func GetRank(c *gin.Context) {
	updateRank := c.Query("update_rank") == "true"
	_page := c.Query("page")

	page, err := strconv.Atoi(_page)

	if err != nil && _page != "" {
		restErr := config.NewBadRequestErr("page must be integer")
		c.JSON(restErr.Code, restErr)
		return
	}

	_perPage := c.Query("perPage")

	perPage, err := strconv.Atoi(_perPage)

	if err != nil && _perPage != "" {
		restErr := config.NewBadRequestErr("page must be integer")
		c.JSON(restErr.Code, restErr)
		return
	}

	rank, restErr := rankUseCase.GetRank(updateRank, page, perPage)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, rank)
}
