package rankUseCase

import (
	"fmt"
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/database"
	"sync"

	"gorm.io/gorm"
)

func UpdateRank() *config.RestErr {
	db := database.GetDatabase()
	groupByUsers := []models.UserCompletedQuiz{}
	groupByQuiz := []models.UserCompletedQuiz{}

	var wg sync.WaitGroup
	var errGetGroupByUsers error
	var errGetGroupByQuiz error

	wg.Add(2)

	go func() {
		groupByUsers, errGetGroupByUsers = getGroupByUsers(db)
		wg.Done()
	}()
	go func() {
		groupByQuiz, errGetGroupByQuiz = getGroupByQuiz(db)
		wg.Done()
	}()

	wg.Wait()

	if errGetGroupByQuiz != nil || errGetGroupByUsers != nil {
		return config.NewInternaErr("cannot calculate rank")
	}

	for _, u := range groupByUsers {
		userRank := dto.UpdateRankDto{
			UserId: u.UserId,
		}

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
		userRank.BestScoreQuizId = aux.QuizId
		userRank.TotalScore = totalScore

		wg.Add(1)
		go func() {
			updateUserRank(userRank, db)
			wg.Done()
		}()
	}

	wg.Wait()

	return nil
}

func updateUserRank(userRank dto.UpdateRankDto, db *gorm.DB) {
	var existingUserRank models.Rank
	result := db.Where("user_id = ?", userRank.UserId).First(&existingUserRank)

	if !hasRankChanged(&existingUserRank, &userRank) {
		fmt.Println("Rank atualizado")
		return
	}

	if result.RowsAffected == 0 {
		newUserRank := models.Rank{
			TotalScore:      userRank.TotalScore,
			UserId:          userRank.UserId,
			BestScoreQuizId: userRank.BestScoreQuizId,
		}
		db.Create(&newUserRank)
		return
	}

	db.Model(&existingUserRank).
		Where("id = ?", existingUserRank.ID).
		Updates(userRank)
}

func hasRankChanged(a *models.Rank, b *dto.UpdateRankDto) bool {
	return a.UserId != b.UserId ||
		a.BestScoreQuizId != b.BestScoreQuizId ||
		a.TotalScore != b.TotalScore
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
