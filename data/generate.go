package data

import (
	"api/model"
	"fmt"
	weightRand "github.com/mroth/weightedrand"
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
	chooser, _ := weightRand.NewChooser(
		weightRand.Choice{Item: "Complete", Weight: 2},
		weightRand.Choice{Item: "Pending LM/CC", Weight: 4},
		weightRand.Choice{Item: "To Review", Weight: 6},
		weightRand.Choice{Item: "Missing", Weight: 7},
	)
	status := chooser.Pick().(string)
	return status
}

func getRandomPeriod() string {
	quarters := []string{"H1 FY2023", "H2 FY2023", "H2 FY2022", "H1 FY2022"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	quarter := quarters[r.Intn(len(quarters))]
	return quarter
}

func getRandomDiscipline() string {
	disciplines := []string{"Development", "AM", "Testing", "Creative", "Facilities"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return disciplines[r.Intn(len(disciplines))]
}
