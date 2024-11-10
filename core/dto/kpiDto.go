package dto

import "mensina-be/database/models"

type QuizRank struct {
	Quiz   models.Quiz `json:"quiz" gorm:"foreignKey:QuizId"`
	QuizId uint        `json:"quizId"`
	Score  int         `json:"score"`
}

type TagRank struct {
	Tag   models.Tag `json:"tag" gorm:"foreignKey:TagId"`
	TagId uint        `json:"tagId"`
	Score  int         `json:"score"`
}

type UserKpiDto struct {
	UserId      uint          `json:"id"`
	QuizzesRank []QuizRank `json:"quizzesRank"`
	TagsRank    []TagRank  `json:"tagsRank"`
}