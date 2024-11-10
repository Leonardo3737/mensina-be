package migrations

import (
	"mensina-be/database/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Quiz{})
	db.AutoMigrate(models.Answer{})
	db.AutoMigrate(models.Question{})
	db.AutoMigrate(models.UserCompletedQuiz{})
	db.AutoMigrate(models.Tag{})
	db.AutoMigrate(models.Rank{})
}
