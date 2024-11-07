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
	Score     int
	UserId    uint
	Total     int
	QuizzId   uint
	Questions map[int]Status
}
