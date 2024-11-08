package tagController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/tagUseCase"

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
		restErr := config.NewInternaErr("cannot list tags")
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(200, tags)
}
