package controller

import (
	"api/data"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func ShowFeedbacks(c *gin.Context) {
	discipline := strings.ToLower(c.Query("discipline"))
	period := strings.ToLower(c.Query("period"))
	status := strings.ToLower(c.Query("status"))
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.Query("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	filteredFeedbacks := data.Feedbacks

	if discipline != "" {
		filteredFeedbacks = filterByDiscipline(filteredFeedbacks, discipline)
	}

	if period != "" {
		filteredFeedbacks = filterByPeriod(filteredFeedbacks, period)
	}

	if status != "" {
		filteredFeedbacks = filterByStatus(filteredFeedbacks, status)
	}

	paginatedFeedbacks := paginateFeedbacks(filteredFeedbacks, limit, offset)

	// Return response
	// return it
	c.JSON(200, gin.H{
		"feedbacks": paginatedFeedbacks,
		"offset":    offset,
		"limit":     len(paginatedFeedbacks),
		"total":     len(filteredFeedbacks),
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func StatusSummary(c *gin.Context) {
	statusCounts := make(map[string]int)

	for _, f := range data.Feedbacks {
		statusCounts[f.Status]++
	}

	summary := make([]map[string]interface{}, 0)
	for status, count := range statusCounts {
		summary = append(summary, map[string]interface{}{
			"status": strings.Title(status),
			"total":  count,
		})
	}

	totalFeedbacks := len(data.Feedbacks)

	response := map[string]interface{}{
		"summary": summary,
		"total":   totalFeedbacks,
	}

	c.JSON(200, response)
}
