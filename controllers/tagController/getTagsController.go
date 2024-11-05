package tagController

import (
	"mensina-be/core/useCases/tagUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get all tags
// @Tags Tag
// @Produce json
// @Success 200 {object} models.Tag "Success"
// @Security BearerAuth
// @Router /tag [get]
func GetTags(c *gin.Context) {
	tags, err := tagUseCase.GetTags()

	if err != nil {
		c.JSON(500, utils.ErrorResponse{
			Error: "internal server error",
		})
		return
	}
	c.JSON(200, tags)
}
