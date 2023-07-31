package controller

import (
	"api/data"
	"api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
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
	search := c.Query("search")
	searchQuery, err := url.QueryUnescape(search)

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	filteredFeedbacks := data.Feedbacks

	if searchQuery != "" {
		filteredFeedbacks = searchFeedbacks(filteredFeedbacks, strings.ToLower(searchQuery))
	}

	if discipline != "" {
		filteredFeedbacks = filterByDiscipline(filteredFeedbacks, discipline)
	}

	if status != "" {
		filteredFeedbacks = filterByStatus(filteredFeedbacks, status)
	}

	if period != "" {
		filteredFeedbacks = filterByPeriod(filteredFeedbacks, period)
	}

	paginatedFeedbacks := paginateFeedbacks(filteredFeedbacks, limit, offset)

	c.JSON(200, gin.H{
		"feedbacks": paginatedFeedbacks,
		"offset":    offset,
		"limit":     len(paginatedFeedbacks),
		"total":     len(filteredFeedbacks),
	})
}

func searchFeedbacks(feedbacks []model.Feedbacks, query string) []model.Feedbacks {
	var results []model.Feedbacks
	searchQuery := strings.ReplaceAll(query, "+", " ")
	for _, feedback := range feedbacks {
		if strings.Contains(strings.ToLower(feedback.Name), searchQuery) {
			results = append(results, feedback)
		}
	}
	return results
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

	// 1. Parse the query parameters
	discipline := strings.ToLower(c.Query("discipline"))
	period := strings.ToLower(c.Query("period"))

	// 2. Filter the data based on the provided discipline and period
	filteredFeedbacks := filterFeedbacks(data.Feedbacks, discipline, period)

	statusCounts := make(map[string]int)

	for _, f := range filteredFeedbacks {
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

	totalFeedbacks := len(filteredFeedbacks)

	response := map[string]interface{}{
		"summary": summary,
		"total":   totalFeedbacks,
	}

	c.JSON(200, response)
}

func SelectOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
	c.Header("Access-Control-Expose-Headers", "Content-Length")
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

func UpdateComment(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
	c.Header("Access-Control-Expose-Headers", "Content-Length")
	id := c.Param("id")

	var update struct {
		Comment string `json:"comments"`
		Status  string `json:"status"`
	}
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var feedback *model.Feedbacks
	for i, fb := range data.Feedbacks {
		if fb.ID == id {
			feedback = &data.Feedbacks[i]
			break
		}
	}

	if feedback == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	if update.Comment != "" {
		feedback.Comments = update.Comment
	}

	validStatus := map[string]bool{
		"Complete":      true,
		"Pending LM/CC": true,
		"To Review":     true,
		"Missing":       true,
	}

	if update.Status != "" {
		if !validStatus[update.Status] {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid status"})
			return
		}
		feedback.Status = update.Status
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback updated", "data": feedback})
}
