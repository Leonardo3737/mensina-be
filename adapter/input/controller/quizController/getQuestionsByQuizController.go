package quizController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/quizUseCase"
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
		restErr := config.NewBadRequestErr("ID has to be integer")
		c.JSON(restErr.Code, restErr)
		return
	}
	questions, err := quizUseCase.GetQuestionByQuiz(id)

	if err != nil {
		restErr := config.NewNotFoundErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(200, questions)
}
