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
			LineManager:          getRandomManager(),
			CareerCoach:          fmt.Sprintf("Coach %d", i),
			LineManagerDone:      getRandomDone(),
			CareerCoachDone:      getRandomDone(),
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

func getRandomDone() bool {
	chooser, _ := weightRand.NewChooser(
		weightRand.Choice{Item: true, Weight: 6},
		weightRand.Choice{Item: false, Weight: 4},
	)
	done := chooser.Pick().(bool)
	return done
}

func getRandomPeriod() string {
	chooser, _ := weightRand.NewChooser(
		weightRand.Choice{Item: "H1 FY2023", Weight: 3},
		weightRand.Choice{Item: "H2 FY2023", Weight: 1},
		weightRand.Choice{Item: "H2 FY2022", Weight: 4},
		weightRand.Choice{Item: "H1 FY2022", Weight: 6},
	)
	period := chooser.Pick().(string)
	return period
}

func getRandomDiscipline() string {
	disciplines := []string{"Development", "AM", "Testing", "Creative", "Facilities"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return disciplines[r.Intn(len(disciplines))]
}

func getRandomManager() string {
	chooser, _ := weightRand.NewChooser(
		weightRand.Choice{Item: "Sophie", Weight: 2},
		weightRand.Choice{Item: "Liam", Weight: 2},
		weightRand.Choice{Item: "Olivia", Weight: 2},
		weightRand.Choice{Item: "Noah", Weight: 2},
		weightRand.Choice{Item: "Emma", Weight: 2},
		weightRand.Choice{Item: "Jackson", Weight: 2},
		weightRand.Choice{Item: "Ava", Weight: 2},
		weightRand.Choice{Item: "Lucas", Weight: 2},
		weightRand.Choice{Item: "Isabella", Weight: 2},
		weightRand.Choice{Item: "Aiden", Weight: 2},
		weightRand.Choice{Item: "Mia", Weight: 2},
		weightRand.Choice{Item: "Caden", Weight: 2},
		weightRand.Choice{Item: "Charlotte", Weight: 2},
		weightRand.Choice{Item: "Mason", Weight: 2},
		weightRand.Choice{Item: "Amelia", Weight: 2},
		weightRand.Choice{Item: "Oliver", Weight: 2},
		weightRand.Choice{Item: "Harper", Weight: 2},
		weightRand.Choice{Item: "Elijah", Weight: 2},
		weightRand.Choice{Item: "Evelyn", Weight: 2},
		weightRand.Choice{Item: "Grayson", Weight: 2},
		weightRand.Choice{Item: "Abigail", Weight: 2},
		weightRand.Choice{Item: "Jacob", Weight: 2},
		weightRand.Choice{Item: "Emily", Weight: 2},
		weightRand.Choice{Item: "Michael", Weight: 2},
		weightRand.Choice{Item: "Elizabeth", Weight: 2},
		weightRand.Choice{Item: "Benjamin", Weight: 2},
		weightRand.Choice{Item: "Ella", Weight: 2},
		weightRand.Choice{Item: "Carter", Weight: 2},
		weightRand.Choice{Item: "Sofia", Weight: 2},
		weightRand.Choice{Item: "James", Weight: 2},
		weightRand.Choice{Item: "Avery", Weight: 2},
		weightRand.Choice{Item: "Alexander", Weight: 2},
		weightRand.Choice{Item: "Grace", Weight: 2},
		weightRand.Choice{Item: "Matthew", Weight: 2},
		weightRand.Choice{Item: "Chloe", Weight: 2},
		weightRand.Choice{Item: "Wyatt", Weight: 2},
		weightRand.Choice{Item: "Scarlett", Weight: 2},
		weightRand.Choice{Item: "Luke", Weight: 2},
		weightRand.Choice{Item: "Victoria", Weight: 2},
		weightRand.Choice{Item: "Henry", Weight: 2},
		weightRand.Choice{Item: "Luna", Weight: 2},
		weightRand.Choice{Item: "Andrew", Weight: 2},
		weightRand.Choice{Item: "Penelope", Weight: 2},
		weightRand.Choice{Item: "Daniel", Weight: 2},
		weightRand.Choice{Item: "Layla", Weight: 2},
		weightRand.Choice{Item: "Logan", Weight: 2},
		weightRand.Choice{Item: "Zoey", Weight: 2},
		weightRand.Choice{Item: "William", Weight: 2},
		weightRand.Choice{Item: "Nora", Weight: 2},
		weightRand.Choice{Item: "Gabriel", Weight: 2},
	)
	manager := chooser.Pick().(string)
	return manager
}
