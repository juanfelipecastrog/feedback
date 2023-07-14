package controller

import (
	"api/data"
	"api/model"
	"github.com/gin-gonic/gin"
	"math"
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

func Search(c *gin.Context) {
	escapedName := c.Query("q")
	name, err := url.QueryUnescape(escapedName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'q' parameter",
		})
		return
	}

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'q' parameter",
		})
		return
	}

	name = strings.ReplaceAll(name, "+", " ")

	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}
	limit := 10

	var matches []model.Feedbacks

	for _, feedback := range data.Feedbacks {
		if strings.Contains(strings.ToLower(feedback.Name), strings.ToLower(name)) {
			matches = append(matches, feedback)
		}
	}

	totalResults := len(matches)
	totalPages := int(math.Ceil(float64(totalResults) / float64(limit)))

	if page > totalPages {
		page = totalPages
	}

	var results []model.Feedbacks
	if len(matches) > 0 {
		startIndex := (page - 1) * limit
		endIndex := int(math.Min(float64(page*limit), float64(totalResults)))
		results = matches[startIndex:endIndex]
	}

	response := gin.H{
		"search": gin.H{
			"totalResults": totalResults,
			"currentPage":  page,
			"pageSize":     limit,
			"totalPages":   totalPages,
			"results":      results,
		},
	}

	c.JSON(http.StatusOK, response)
}

func UpdateComment(c *gin.Context) {
	id := c.Param("id")

	var update struct {
		Comment string `json:"comment"`
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

	if update.Status != "" {
		feedback.Status = update.Status
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback updated", "data": feedback})
}
