package userUseCase

import (
	"mensina-be/config"
	"mensina-be/core/dto"
	"mensina-be/database"
	"mensina-be/database/models"
	"sync"

	"gorm.io/gorm"
)

func GetUserKpi(userId uint) (dto.UserKpiDto, *config.RestErr) {
	db := database.GetDatabase()
	var wg sync.WaitGroup
	var userKpi dto.UserKpiDto

	wg.Add(3)

	go func() {
		defer wg.Done()

		getQuizRank(userId, db, &userKpi.QuizzesRank)
	}()
	go func() {
		defer wg.Done()

		getAvarage(userId, db, &userKpi)
	}()
	go func() {
		defer wg.Done()

		getTagRank(userId, db, &userKpi.TagsRank)

		for _, tagRank := range userKpi.TagsRank {
			userKpi.TotalScore += tagRank.TotalScore
		}
	}()

	wg.Wait()

	return userKpi, nil
}

func getQuizRank(userId uint, db *gorm.DB, quizRank *[]dto.QuizRank) {
	db.
		Model(&models.UserCompletedQuiz{}).
		Select(`
		quiz_id, 
		q.title as quiz_title, 
		q.tag_id as tag_id, 
		t.description as tag_description, 
		MAX(score) AS score`).
		InnerJoins("INNER JOIN quizzes q on q.id = user_completed_quizzes.quiz_id").
		InnerJoins("INNER JOIN tags t on t.id = q.tag_id").
		Where("user_id = ?", userId).
		Group("quiz_id").
		Order("score DESC").
		Find(&quizRank)
}

func getTagRank(userId uint, db *gorm.DB, tagRank *[]dto.TagRank) {
	subQueryTotalScore := db.
		Model(&models.UserCompletedQuiz{}).
		Select("quiz_id, MAX(score) as score").
		Where("user_id = ?", userId).
		Group("quiz_id")

	db.
		Model(&models.Tag{}).
		Select(`
		tags.id as tag_id, 
		tags.description as tag_description, 
		SUM(sq.score) as total_score`).
		Joins("INNER JOIN quizzes q ON q.tag_id = tags.id").
		Joins("INNER JOIN (?) AS sq ON sq.quiz_id = q.id", subQueryTotalScore).
		Group("tags.id, tags.description").
		Order("total_score DESC").
		Scan(tagRank)
}

func getAvarage(userId uint, db *gorm.DB, userKpi *dto.UserKpiDto) {
	db.
		Model(models.UserCompletedQuiz{}).
		Select(`
		user_id,
		SUM(correct_answers)/(COUNT(*)*5) AS correct_answers_avarage
		`).
		Where("user_id = ?", userId).
		Group("user_id").
		Scan(userKpi)
}
