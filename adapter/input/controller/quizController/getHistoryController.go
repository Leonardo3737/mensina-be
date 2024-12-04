package quizController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get quizzes history
// @Tags Quiz
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.UserCompletedQuiz "Success"
// @Router /quiz/history [get]
func GetHistory(c *gin.Context) {
	userId, err := utils.GetUserIdByToken(c)

	if err != nil {
		restErr := config.NewUnauthorizedErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	history, restErr := quizUseCase.GetHistory(userId)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(200, history)
}
