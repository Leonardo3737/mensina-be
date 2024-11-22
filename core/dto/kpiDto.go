package dto

type QuizRank struct {
	QuizTitle      string `json:"quizTitle"`
	QuizId         uint   `json:"quizId"`
	TagDescription string `json:"tagDescription"`
	TagId          uint   `json:"tagId"`
	Score          int    `json:"score"`
}

type TagRank struct {
	TagDescription string `json:"tagDescription"`
	TagId          uint   `json:"tagId"`
	TotalScore     int    `json:"totalscore"`
}

type UserKpiDto struct {
	CorrectAnswersAvarage float64    `json:"correctAnswersAvarage"`
	UserId                uint       `json:"userId"`
	TotalScore            int        `json:"totaScore"`
	QuizzesRank           []QuizRank `json:"quizzesRank"`
	TagsRank              []TagRank  `json:"tagsRank"`
}
