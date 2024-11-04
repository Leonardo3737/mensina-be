package quizController

import (
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get question by Quiz_ID
// @Tags Quiz
// @Param quiz_id path int true "QUIZ ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.OutputQuestionDto "Success"
// @Router /quiz/questions/{quiz_id} [get]
func GetQuestionByQuiz(c *gin.Context) {
	_id := c.Param("quiz_id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "ID has to be integer",
		})
		return
	}
	questions, err := quizUseCase.GetQuestionByQuiz(id)

	if err != nil {
		c.JSON(404, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, questions)
}
