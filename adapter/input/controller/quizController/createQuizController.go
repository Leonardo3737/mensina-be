package quizController

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/useCases/quizUseCase"

	"github.com/gin-gonic/gin"
)

// @Summary Create quiz
// @Tags Quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param quiz body dto.CreateQuizDto true "Quiz data"
// @Success 201 {object} models.Quiz "Quiz"
// @Router /quiz [post]
func CreateQuiz(c *gin.Context) {
	var _Quiz dto.CreateQuizDto

	if err := c.ShouldBindJSON(&_Quiz); err != nil {
		restErr := config.NewBadRequestErr("some field are incorrect")
		c.JSON(restErr.Code, restErr)
		return
	}

	newQuiz, restErr := quizUseCase.CreateQuiz(&_Quiz)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(201, newQuiz)
}
