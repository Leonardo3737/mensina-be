package quizController

import (
	"mensina-be/config"
	"mensina-be/core/routines"
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Start quiz
// @Tags Quiz
// @Success 200 {object} dto.QuizSession "Success"
// @Param quiz_id path string true "Quiz ID"
// @Security BearerAuth
// @Router /quiz/start/{quiz_id} [get]
func StartQuiz(c *gin.Context, quizRoutineChannel chan routines.RoutineCallback) {
	userId, err := utils.GetUserIdByToken(c)

	if err != nil {
		restErr := config.NewUnauthorizedErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	quizIdStr := c.Param("quiz_id")
	quizId, err := strconv.Atoi(quizIdStr)

	if err != nil {
		restErr := config.NewBadRequestErr("Invalid quiz ID")
		c.JSON(restErr.Code, restErr)
		return
	}

	quizSession, restErr := quizUseCase.StartQuiz(userId, uint(quizId), quizRoutineChannel)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, quizSession)
}
