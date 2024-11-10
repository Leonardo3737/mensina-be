package models

type Rank struct {
	ID              uint `json:"id" gorm:"primaryKey"`
	TotalScore      int  `json:"totalScore"`
	UserId          uint `json:"userId"`
	User            User `json:"user" gorm:"foreignKey:UserId"`
	BestScoreQuizId uint `json:"bestScoreQuizId"`
	BestScoreQuiz   Quiz `json:"bestScoreQuiz" gorm:"foreignKey:BestScoreQuizId"`
}
