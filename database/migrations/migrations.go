package migrations

import (
	"mensina-be/core/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Quiz{})
	db.AutoMigrate(models.Answer{})
	db.AutoMigrate(models.Question{})
	db.AutoMigrate(models.UserCompletedQuiz{})
	db.AutoMigrate(models.Tag{})
}
