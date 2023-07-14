package controller

import (
	"api/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func ShowFeedbacks(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
	c.Header("Access-Control-Expose-Headers", "Content-Length")
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

	c.JSON(200, gin.H{
		"feedbacks": paginatedFeedbacks,
		"offset":    offset,
		"limit":     len(paginatedFeedbacks),
		"total":     len(filteredFeedbacks),
	})
}

func HealthCheck(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
	c.Header("Access-Control-Expose-Headers", "Content-Length")
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func StatusSummary(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
	c.Header("Access-Control-Expose-Headers", "Content-Length")
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

	sort.Slice(summary, func(i, j int) bool {
		order := map[string]int{
			"Complete":      0,
			"To Review":     1,
			"Pending LM/CC": 2,
			"Missing":       3,
		}
		return order[summary[i]["status"].(string)] < order[summary[j]["status"].(string)]
	})

	totalFeedbacks := len(data.Feedbacks)

	response := map[string]interface{}{
		"summary": summary,
		"total":   totalFeedbacks,
	}

	c.JSON(200, response)
}

func SelectOptions(c *gin.Context) {
	options := make(map[string][]string)
	options["disciplines"] = []string{
		"Analysis",
		"Application Management",
		"Business Platform",
		"Development",
		"Data",
		"Project Delivery Management",
		"Testing",
	}
	options["status"] = []string{"Complete", "Pending LM/CC", "To Review", "Missing"}
	options["periods"] = []string{"H1 FY2023", "H2 FY2023", "H2 FY2022", "H1 FY2022"}

	c.JSON(http.StatusOK, options)

}
