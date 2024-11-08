package dto

type RankDto struct {
	TotalScore         int    `json:"totalScore"`
	BestScoreQuizId    uint   `json:"bestScoreQuizId"`
	UserId             uint   `json:"userId"`
	Username           string `json:"username"`
	BestScoreQuizTitle string `json:"bestScoreQuizTitle"`
}
