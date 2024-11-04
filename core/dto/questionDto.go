package dto

type OutputQuestionDto struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Answers     []OutputAnswerDto `json:"answers"`
}
