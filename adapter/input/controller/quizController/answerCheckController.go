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

type answerCheckResponse struct {
	IsCorrect bool `json:"is_correct"`
}

// @Summary Check
// @Tags Quiz
// @Produce json
// @Success 200 {object} quizController.answerCheckResponse "Success"
// @Param answer_id query string true "Answer ID"
// @Param question_id query string true "Question ID"
// @Security BearerAuth
// @Router /quiz/answer_check [get]
func AnswerCheck(c *gin.Context, quizRoutineChannel chan routines.RoutineCallback) {
	_answerId := c.Query("answer_id")
	_questionId := c.Query("question_id")

	answerId, err := strconv.Atoi(_answerId)

	if err != nil {
		restErr := config.NewBadRequestErr("answer_id must be integer")
		c.JSON(restErr.Code, restErr)
		return
	}

	questionId, err := strconv.Atoi(_questionId)

	if err != nil {
		restErr := config.NewBadRequestErr("question_id must be integer")
		c.JSON(restErr.Code, restErr)
		return
	}

	userId, err := utils.GetUserIdByToken(c)

	if err != nil {
		restErr := config.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	isCorrect, restErr := quizUseCase.AnswerCheck(answerId, questionId, int(userId), quizRoutineChannel)

	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, answerCheckResponse{
		IsCorrect: isCorrect,
	})
}
