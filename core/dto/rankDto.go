package dto

type RankDto struct {
	TotalScore         int    `json:"totalScore"`
	BestScoreQuizId    uint   `json:"bestScoreQuizId"`
	UserId             uint   `json:"userId"`
	Username           string `json:"username"`
	BestScoreQuizTitle string `json:"bestScoreQuizTitle"`
}

type UpdateRankDto struct {
	TotalScore      int  `json:"totalScore"`
	UserId          uint `json:"userId"`
	BestScoreQuizId uint `json:"bestScoreQuizId"`
}
