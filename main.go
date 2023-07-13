package main

import (
	"api/controller"
	"api/data"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	data.Feedbacks = data.GenerateFeedbacks(5000)
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health-check/", controller.HealthCheck)
	r.GET("/feedback/", controller.ShowFeedbacks)
	r.GET("/status/", controller.StatusSummary)
	r.Run() // listen and serve on 0.0.0.0:8080
}
