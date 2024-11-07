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

		/* sectionCh := <-ch
		section, exist := Sections[int(sectionCh.UserId)]

		if !exist {
			fmt.Printf("Iniciando Quiz, userId: %d | quizId: %d\n", sectionCh.UserId, sectionCh.QuizzId)
			Sections[int(sectionCh.UserId)] = &dto.QuizSession{
				Total:   0,
				Score:   0,
				QuizzId: sectionCh.QuizzId,
			}
			continue
		}

		section.Total++
		section.Score = +sectionCh.Score

		fmt.Printf("sectionCh.Score: %x\n", sectionCh.Score)

		fmt.Printf("Total: %x\n", section.Total)
		fmt.Printf("Respondidos: %x\n", section.Score)
		if section.Total == 5 {
			fmt.Println("Finalizando Quiz")
			ch <- *section
			fmt.Println("passou")
			delete(Sections, int(sectionCh.UserId))
		} */
	}
}
