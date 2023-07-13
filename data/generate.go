package data

import (
	"api/model"
	"fmt"
	"math/rand"
	"time"
)

var Feedbacks []model.Feedbacks

func GenerateFeedbacks(quantity int) []model.Feedbacks {

	for i := 1; i <= quantity; i++ {
		feedback := model.Feedbacks{
			ID:                   fmt.Sprintf("%d", i),
			Name:                 fmt.Sprintf("Employee %d", i),
			LineManager:          fmt.Sprintf("Manager %d", i),
			CareerCoach:          fmt.Sprintf("Coach %d", i),
			Authorized:           fmt.Sprintf("Authorized %d", i),
			CompensationReviewer: fmt.Sprintf("Reviewer %d", i),
			Comments:             fmt.Sprintf("Feedback %d", i),
			Status:               getRandomStatus(),
			Discipline:           getRandomDiscipline(),
			Period:               getRandomPeriod(),
		}
		Feedbacks = append(Feedbacks, feedback)
	}

	return Feedbacks
}

func getRandomStatus() string {
	statuses := []string{"Complete", "Pending LM/CC", "To Review", "Missing"}
	rand.Seed(time.Now().UnixNano())
	return statuses[rand.Intn(len(statuses))]
}

func getRandomPeriod() string {
	quarters := []string{"H1 FY2023", "H2 FY2023", "H2 FY2022", "H1 FY2022"}
	rand.Seed(time.Now().UnixNano())
	quarter := quarters[rand.Intn(len(quarters))]
	return quarter
}

func getRandomDiscipline() string {
	disciplines := []string{"Development", "AM", "Testing", "Creative", "Facilities"}
	rand.Seed(time.Now().UnixNano())
	return disciplines[rand.Intn(len(disciplines))]
}
