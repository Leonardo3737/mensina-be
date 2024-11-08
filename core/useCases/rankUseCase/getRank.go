package rankUseCase

import (
	"mensina-be/core/dto"
	"mensina-be/core/models"
	"mensina-be/database"
	"sync"

	"gorm.io/gorm"
)

func GetRank() ([]dto.RankDto, int, error) {
	db := database.GetDatabase()
	groupByUsers := []models.UserCompletedQuiz{}
	groupByQuiz := []models.UserCompletedQuiz{}

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		groupByUsers = getGroupByUsers(db)
	}()
	go func() {
		defer wg.Done()
		groupByQuiz = getGroupByQuiz(db)
	}()

	wg.Wait()

	var rank []dto.RankDto

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

	return rank, 200, nil
}

func getGroupByUsers(db *gorm.DB) []models.UserCompletedQuiz {
	var groupByUsers []models.UserCompletedQuiz
	db.
		Model(&models.UserCompletedQuiz{}).
		Select("user_id").
		Group("user_id").
		Preload("User").
		Where("score > ?", 0).
		Find(&groupByUsers)

	return groupByUsers
}

func getGroupByQuiz(db *gorm.DB) []models.UserCompletedQuiz {
	var groupByQuiz []models.UserCompletedQuiz
	db.
		Model(&models.UserCompletedQuiz{}).
		Select("quiz_id, user_id, MAX(score) as score").
		Group("quiz_id, user_id").
		Order("score ASC").
		Preload("Quiz").
		Find(&groupByQuiz)

	return groupByQuiz
}
