package model

type Feedback struct {
	Feedbacks []Feedbacks `json:"feedbacks"`
	Offset    int         `json:"offset"`
	Limit     int         `json:"limit"`
	Total     int         `json:"total"`
}

type Feedbacks struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	LineManager          string `json:"line_manager"`
	CareerCoach          string `json:"career_coach"`
	Authorized           string `json:"authorized"`
	CompensationReviewer string `json:"compensation_reviewer"`
	Comments             string `json:"comments"`
	Status               string `json:"status"`
	Discipline           string `json:"discipline"`
	Period               string `json:"period"`
}
