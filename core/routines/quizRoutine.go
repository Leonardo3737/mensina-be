package routines

import (
	"fmt"
	"mensina-be/core/dto"
)

func RunQuizRoutine(ch chan dto.QuizRoutineChannel) {
	var Sections = make(map[int]*dto.QuizRoutineChannel)
	for {
		sectionCh := <-ch
		section, exist := Sections[int(sectionCh.UserId)]

		if !exist {
			fmt.Printf("Iniciando Quiz, userId: %d | quizId: %d\n", sectionCh.UserId, sectionCh.QuizzId)
			Sections[int(sectionCh.UserId)] = &dto.QuizRoutineChannel{
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
		}
	}
}
