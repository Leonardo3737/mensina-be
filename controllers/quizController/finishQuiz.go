package quizController

import (
	"mensina-be/core/routines"
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Finish quiz
// @Tags Quiz
// @Success 204 "Success"
// @Param quiz_id path string true "Quiz ID"
// @Security BearerAuth
// @Router /quiz/finish/{quiz_id} [delete]
func FinishQuiz(c *gin.Context, ch chan routines.RoutineCallback) {
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
	quizUseCase.FinishQuiz(uint(quizId), userId, ch)
	c.Status(204)
}
