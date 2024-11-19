package dto

type Status int8

const (
	Unanswered Status = iota
	InCorrect
	Correct
)

func (s Status) String() string {
	switch s {
	case Correct:
		return "Correct"
	case InCorrect:
		return "InCorrect"
	default:
		return "Unanswered"
	}
}

type QuizSession struct {
	Score     int            `json:"score"`
	Total     int            `json:"total"`
	Correct   int            `json:"correct"`
	UserId    uint           `json:"userId"`
	QuizzId   uint           `json:"quizzId"`
	QuizTitle string         `json:"quizTitle"`
	Questions map[int]Status `json:"questions"`
}
