package quizController

import (
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get all quizzes
// @Tags Quiz
// @Produce json
// @Success 200 {array} models.Quiz "Success"
// @Security BearerAuth
// @Router /quiz [get]
func GetQuiz(c *gin.Context) {
	quizzes, err := quizUseCase.GetQuizzes()

	if err != nil {
		c.JSON(500, utils.ErrorResponse{
			Error: "cannot list users",
		})
		return
	}

	c.JSON(200, quizzes)
}
