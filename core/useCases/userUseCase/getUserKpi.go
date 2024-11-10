package userUseCase

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"

	"gorm.io/gorm"
)

type dataChannel struct {
	err  error
	data []models.UserCompletedQuiz
}

func GetUserKpi(userId uint) (dto.UserKpiDto, *config.RestErr) {
	db := database.GetDatabase()
	ch := make(chan dataChannel)

	go getUserCompletedQuiz(userId, db, ch)

	user, userErr := GetUserInfos(userId)
	dataChannel := <-ch
	if dataChannel.err != nil || userErr != nil {
		return dto.UserKpiDto{}, config.NewInternaErr("cannot calculate kpi")
	}
	userKpi := dto.UserKpiDto{
		UserId: user.ID,
		QuizzesRank: make([]dto.QuizRank, 0, len(dataChannel.data)),
		TagsRank: make([]dto.TagRank, 0, len(dataChannel.data)),
	}
	
}

func getUserCompletedQuiz(userId uint, db *gorm.DB, ch chan dataChannel) {
	var userCompletedQuiz []models.UserCompletedQuiz 
	err := db.
	Preload("Quiz").
	Where("user_id = ?", userId).
	Order("score DESC").
	Find(&userCompletedQuiz).Error

	dataChannel := dataChannel{
		err: err,
		data: userCompletedQuiz,
	}
	ch <- dataChannel
}