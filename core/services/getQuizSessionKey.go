package services

import "fmt"

func GetQuizSessionsKey(userId, quizId uint) string {
	return fmt.Sprintf("%x - %x", userId, quizId)
}
