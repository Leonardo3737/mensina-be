package tagController

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/useCases/tagUseCase"

	"github.com/gin-gonic/gin"
)

// @Summary Create tag
// @Tags Tag
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body dto.CreateTagDto true "Tag data"
// @Success 201 {object} models.Tag "Tag"
// @Router /tag [post]
func CreateTag(c *gin.Context) {
	var _Tag dto.CreateTagDto

	if err := c.ShouldBindJSON(&_Tag); err != nil {
		restErr := config.NewBadRequestErr("some field are incorrect")
		c.JSON(restErr.Code, restErr)
		return
	}

	newTag, restErr := tagUseCase.CreateTag(&_Tag)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(201, newTag)
}
