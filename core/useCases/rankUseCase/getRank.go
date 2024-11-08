package rankUseCase

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/database"
	"sync"

	"gorm.io/gorm"
)

func GetRank() ([]dto.RankDto, *config.RestErr) {
	db := database.GetDatabase()
	groupByUsers := []models.UserCompletedQuiz{}
	groupByQuiz := []models.UserCompletedQuiz{}

	var wg sync.WaitGroup
	var errGetGroupByUsers error
	var errGetGroupByQuiz error
	wg.Add(2)
	go func() {
		defer wg.Done()
		groupByUsers, errGetGroupByUsers = getGroupByUsers(db)
	}()
	go func() {
		defer wg.Done()
		groupByQuiz, errGetGroupByQuiz = getGroupByQuiz(db)
	}()

	wg.Wait()

	var rank []dto.RankDto

	if errGetGroupByQuiz != nil || errGetGroupByUsers != nil {
		return rank, config.NewInternaErr("cannot calculate rank")
	}

	for i, u := range groupByUsers {
		rank = append(rank, dto.RankDto{
			UserId:   u.UserId,
			Username: u.User.UserName,
		})

		aux := models.UserCompletedQuiz{}
		totalScore := 0

		for _, q := range groupByQuiz {
			if q.UserId == u.UserId {
				totalScore += q.Score
				if q.Score > aux.Score {
					aux = q
				}
			}
		}
		rank[i].BestScoreQuizId = aux.QuizId
		rank[i].BestScoreQuizTitle = aux.Quiz.Title
		rank[i].TotalScore = totalScore
	}

	return rank, nil
}

func getGroupByUsers(db *gorm.DB) ([]models.UserCompletedQuiz, error) {
	var groupByUsers []models.UserCompletedQuiz
	err := db.
		Model(&models.UserCompletedQuiz{}).
		Select("user_id").
		Group("user_id").
		Preload("User").
		Where("score > ?", 0).
		Find(&groupByUsers).Error

	return groupByUsers, err
}

func getGroupByQuiz(db *gorm.DB) ([]models.UserCompletedQuiz, error) {
	var groupByQuiz []models.UserCompletedQuiz
	err := db.
		Model(&models.UserCompletedQuiz{}).
		Select("quiz_id, user_id, MAX(score) as score").
		Group("quiz_id, user_id").
		Order("score ASC").
		Preload("Quiz").
		Find(&groupByQuiz).Error

	return groupByQuiz, err
}
