package data

import (
	"api/model"
	"fmt"
	weightRand "github.com/mroth/weightedrand"
	"math/rand"
	"time"
)

var Feedbacks []model.Feedbacks

func GenerateFeedbacks(employees int) []model.Feedbacks {

	for i := 1; i <= employees; i++ {
		feedback := model.Feedbacks{
			ID:                   fmt.Sprintf("%d", i),
			Name:                 getRandomName(),
			LineManager:          getRandomManager(),
			CareerCoach:          getRandomCoach(),
			LineManagerDone:      getRandomDone(),
			CareerCoachDone:      getRandomDone(),
			Authorized:           getRandomManager(),
			CompensationReviewer: getRandomCoach(),
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

func getRandomCoach() string {
	chooser, _ := weightRand.NewChooser(
		weightRand.Choice{Item: "Sophia", Weight: 2},
		weightRand.Choice{Item: "Liam", Weight: 2},
		weightRand.Choice{Item: "Isabella", Weight: 2},
		weightRand.Choice{Item: "Noah", Weight: 2},
		weightRand.Choice{Item: "Emma", Weight: 2},
		weightRand.Choice{Item: "Jackson", Weight: 2},
		weightRand.Choice{Item: "Ava", Weight: 2},
		weightRand.Choice{Item: "Lucas", Weight: 2},
		weightRand.Choice{Item: "Mia", Weight: 2},
		weightRand.Choice{Item: "Elijah", Weight: 2},
		weightRand.Choice{Item: "Charlotte", Weight: 2},
		weightRand.Choice{Item: "Carter", Weight: 2},
		weightRand.Choice{Item: "Amelia", Weight: 2},
		weightRand.Choice{Item: "Olivia", Weight: 2},
		weightRand.Choice{Item: "Benjamin", Weight: 2},
		weightRand.Choice{Item: "Harper", Weight: 2},
		weightRand.Choice{Item: "Alexander", Weight: 2},
		weightRand.Choice{Item: "Avery", Weight: 2},
		weightRand.Choice{Item: "Ethan", Weight: 2},
	)
	coach := chooser.Pick().(string)
	return coach
}

func getRandomName() string {
	chooser, _ := weightRand.NewChooser(
		weightRand.Choice{Item: "Sofia Garcia", Weight: 2},
		weightRand.Choice{Item: "Liam Rodriguez", Weight: 2},
		weightRand.Choice{Item: "Isabella Martínez", Weight: 2},
		weightRand.Choice{Item: "Noah Gonzalez", Weight: 2},
		weightRand.Choice{Item: "Emma Perez", Weight: 2},
		weightRand.Choice{Item: "Lucas Diaz", Weight: 2},
		weightRand.Choice{Item: "Mia Hernandez", Weight: 2},
		weightRand.Choice{Item: "Mateo Lopez", Weight: 2},
		weightRand.Choice{Item: "Valentina Torres", Weight: 2},
		weightRand.Choice{Item: "Emilio Ramírez", Weight: 2},
		weightRand.Choice{Item: "Renata Flores", Weight: 2},
		weightRand.Choice{Item: "Joaquín Sánchez", Weight: 2},
		weightRand.Choice{Item: "Camila Castro", Weight: 2},
		weightRand.Choice{Item: "Sebastian Vargas", Weight: 2},
		weightRand.Choice{Item: "Valeria Guzman", Weight: 2},
		weightRand.Choice{Item: "Maximiliano Ríos", Weight: 2},
		weightRand.Choice{Item: "Isabella Silva", Weight: 2},
		weightRand.Choice{Item: "Benjamin Mendez", Weight: 2},
		weightRand.Choice{Item: "Emilia Figueroa", Weight: 2},
		weightRand.Choice{Item: "Tomas Rojas", Weight: 2},
		weightRand.Choice{Item: "Valentina Reyes", Weight: 2},
		weightRand.Choice{Item: "Juan Pablo Cruz", Weight: 2},
		weightRand.Choice{Item: "Gabriela Morales", Weight: 2},
		weightRand.Choice{Item: "Diego Herrera", Weight: 2},
		weightRand.Choice{Item: "Mariana Ortega", Weight: 2},
		weightRand.Choice{Item: "Javier Miranda", Weight: 2},
		weightRand.Choice{Item: "Andrea Paredes", Weight: 2},
		weightRand.Choice{Item: "Carlos Mendoza", Weight: 2},
		weightRand.Choice{Item: "Fernanda Castro", Weight: 2},
		weightRand.Choice{Item: "Alejandro Guzman", Weight: 2},
		weightRand.Choice{Item: "Paula Soto", Weight: 2},
		weightRand.Choice{Item: "Miguel Ángel León", Weight: 2},
		weightRand.Choice{Item: "Catalina Medina", Weight: 2},
		weightRand.Choice{Item: "Felipe Gallegos", Weight: 2},
		weightRand.Choice{Item: "Daniela Vargas", Weight: 2},
		weightRand.Choice{Item: "Lucia Hernandez", Weight: 2},
		weightRand.Choice{Item: "Rie Kaneko", Weight: 2},
		weightRand.Choice{Item: "Thomas Fauquemberg", Weight: 2},
		weightRand.Choice{Item: "Juan Pablo Rodriguez", Weight: 2},
		weightRand.Choice{Item: "Valentina García", Weight: 2},
		weightRand.Choice{Item: "Martin Torres", Weight: 2},
		weightRand.Choice{Item: "Renata Morales", Weight: 2},
		weightRand.Choice{Item: "Javier López", Weight: 2},
		weightRand.Choice{Item: "Florencia Martinez", Weight: 2},
		weightRand.Choice{Item: "Emilio Silva", Weight: 2},
		weightRand.Choice{Item: "Camila Vargas", Weight: 2},
		weightRand.Choice{Item: "Mariano Castro", Weight: 2},
		weightRand.Choice{Item: "Isabella Ramirez", Weight: 2},
		weightRand.Choice{Item: "Diego Guzman", Weight: 2},
		weightRand.Choice{Item: "Sofia Rios", Weight: 2},
		weightRand.Choice{Item: "Joaquin Soto", Weight: 2},
		weightRand.Choice{Item: "Valeria Flores", Weight: 2},
		weightRand.Choice{Item: "Matias Herrera", Weight: 2},
		weightRand.Choice{Item: "Ana Paula Ortega", Weight: 2},
		weightRand.Choice{Item: "Leonardo Paredes", Weight: 2},
		weightRand.Choice{Item: "Antonella Diaz", Weight: 2},
		weightRand.Choice{Item: "Angel Medina", Weight: 2},
		weightRand.Choice{Item: "Maria Jose Cruz", Weight: 2},
		weightRand.Choice{Item: "Emmanuel Gallegos", Weight: 2},
		weightRand.Choice{Item: "Valentina Mendoza", Weight: 2},
		weightRand.Choice{Item: "Santiago Castro", Weight: 2},
		weightRand.Choice{Item: "Valeria Miranda", Weight: 2},
		weightRand.Choice{Item: "Sebastian Vargas", Weight: 2},
		weightRand.Choice{Item: "Amanda Torres", Weight: 2},
		weightRand.Choice{Item: "Hugo González", Weight: 2},
		weightRand.Choice{Item: "Carolina Flores", Weight: 2},
		weightRand.Choice{Item: "Pablo Soto", Weight: 2},
		weightRand.Choice{Item: "Diego Ramirez", Weight: 2},
		weightRand.Choice{Item: "Valentina Garcia", Weight: 2},
		weightRand.Choice{Item: "Santiago Martínez", Weight: 2},
		weightRand.Choice{Item: "Luciana Morales", Weight: 2},
		weightRand.Choice{Item: "Matias Castro", Weight: 2},
		weightRand.Choice{Item: "Isabella Lopez", Weight: 2},
		weightRand.Choice{Item: "Sebastian Torres", Weight: 2},
		weightRand.Choice{Item: "Valeria Herrera", Weight: 2},
		weightRand.Choice{Item: "Juan Jose Paredes", Weight: 2},
		weightRand.Choice{Item: "Carolina Medina", Weight: 2},
		weightRand.Choice{Item: "Mariano Soto", Weight: 2},
		weightRand.Choice{Item: "Camila Guzman", Weight: 2},
		weightRand.Choice{Item: "Emilio Vargas", Weight: 2},
		weightRand.Choice{Item: "Valentina Silva", Weight: 2},
		weightRand.Choice{Item: "Santiago Ramirez", Weight: 2},
		weightRand.Choice{Item: "Luciana Torres", Weight: 2},
		weightRand.Choice{Item: "Emiliano Mendoza", Weight: 2},
		weightRand.Choice{Item: "Valeria Castro", Weight: 2},
		weightRand.Choice{Item: "Joaquin Herrera", Weight: 2},
		weightRand.Choice{Item: "Catalina Morales", Weight: 2},
		weightRand.Choice{Item: "Miguel Vargas", Weight: 2},
		weightRand.Choice{Item: "Antonella Lopez", Weight: 2},
		weightRand.Choice{Item: "Sebastian Ramirez", Weight: 2},
		weightRand.Choice{Item: "Valentina Soto", Weight: 2},
		weightRand.Choice{Item: "Martin Guzman", Weight: 2},
		weightRand.Choice{Item: "Sofia Paredes", Weight: 2},
		weightRand.Choice{Item: "Mateo Herrera", Weight: 2},
		weightRand.Choice{Item: "Isabella Ramirez", Weight: 2},
		weightRand.Choice{Item: "Lucas Castro", Weight: 2},
		weightRand.Choice{Item: "Emilia Morales", Weight: 2},
		weightRand.Choice{Item: "Maria Jose Cruz", Weight: 2},
		weightRand.Choice{Item: "Mariano Soto", Weight: 2},
		weightRand.Choice{Item: "Emilio Silva", Weight: 2},
	)
	name := chooser.Pick().(string)
	return name
}
