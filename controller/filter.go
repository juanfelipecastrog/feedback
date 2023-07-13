package controller

import (
	"api/model"
	"strings"
)

func filterByDiscipline(feedbacks []model.Feedbacks, discipline string) []model.Feedbacks {
	var filtered []model.Feedbacks
	for _, f := range feedbacks {
		if strings.ToLower(f.Discipline) == strings.ToLower(discipline) {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

func filterByPeriod(feedbacks []model.Feedbacks, period string) []model.Feedbacks {
	var filtered []model.Feedbacks
	for _, f := range feedbacks {
		if strings.EqualFold(strings.ReplaceAll(f.Period, " ", ""), strings.ReplaceAll(period, " ", "")) {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

func filterByStatus(feedbacks []model.Feedbacks, status string) []model.Feedbacks {
	var filtered []model.Feedbacks
	for _, f := range feedbacks {
		if strings.ToLower(f.Status) == status {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

func paginateFeedbacks(feedbacks []model.Feedbacks, limit, offset int) []model.Feedbacks {
	startIndex := offset
	endIndex := offset + limit

	if startIndex < 0 {
		startIndex = 0
	}

	if endIndex > len(feedbacks) {
		endIndex = len(feedbacks)
	}

	return feedbacks[startIndex:endIndex]
}
