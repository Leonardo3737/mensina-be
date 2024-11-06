package quizController

import (
	"mensina-be/core/dto"
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Start quiz
// @Tags Quiz
// @Success 200
// @Param quiz_id path string false "Quiz ID"
// @Security BearerAuth
// @Router /quiz/start/{quiz_id} [get]
func StartQuiz(c *gin.Context, quizRoutineChannel chan dto.QuizRoutineChannel) {
	userId, err := utils.GetUserIdByToken(c)

	if err != nil {
		c.JSON(401, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	quizIdStr := c.Param("quiz_id")
	quizId, err := strconv.Atoi(quizIdStr)
	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "Invalid quiz ID",
		})
		return
	}

	status, err := quizUseCase.StartQuiz(userId, uint(quizId), quizRoutineChannel)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.Status(status)
}
