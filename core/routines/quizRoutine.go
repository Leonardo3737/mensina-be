package routines

import (
	"mensina-be/core/dto"
	"sync"
)

type QuizSessions map[string]*dto.QuizSession
type RoutineCallback func(QuizSessions) *sync.WaitGroup

func RunQuizRoutine(chCallback chan RoutineCallback) {
	var quizSession = make(QuizSessions)
	for {
		callback := <-chCallback
		wg := callback(quizSession)
		if wg != nil {
			wg.Done()
		}
	}
}
