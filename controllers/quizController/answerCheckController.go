package quizController

import (
	"mensina-be/core/routines"
	"mensina-be/core/useCases/quizUseCase"
	"mensina-be/utils"
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
		c.JSON(400, utils.ErrorResponse{
			Error: "answer_id must be integer",
		})
		return
	}

	questionId, err := strconv.Atoi(_questionId)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: "question_id must be integer",
		})
		return
	}

	userId, err := utils.GetUserIdByToken(c)

	if err != nil {
		c.JSON(400, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	isCorrect, status, err := quizUseCase.AnswerCheck(answerId, questionId, int(userId), quizRoutineChannel)

	if err != nil {
		c.JSON(status, utils.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, answerCheckResponse{
		IsCorrect: isCorrect,
	})
}
