package services

import (
	"fmt"
	"strconv"
	"strings"
)

func GetQuizSessionsKey(userId, quizId uint) string {
	return fmt.Sprintf("%x-%x", userId, quizId)
}

func ExtractQuizId(s string) (uint, uint, error) {
	// Divide a string em "userId" e "quizId" usando o separador "-"
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("formato inválido")
	}

	// Converte o quizId de hexadecimal para inteiro
	quizId, err := strconv.ParseInt(parts[1], 16, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("erro ao converter quizId: %v", err)
	}

	userId, err := strconv.ParseInt(parts[0], 16, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("erro ao converter quizId: %v", err)
	}

	return uint(quizId), uint(userId), nil
}
