package main

import (
	"github.com/gin-gonic/gin"

	"quizbot/config"
	"quizbot/internal/api/web"
)

func init() {
	config.LoadConfig()
}

func main() {
	router := gin.Default()
	api := router.Group("/api/")
	q := api.Group("/questions")
	q.POST("/create", web.TimeIt, web.CreateQuestion)
	router.Run(":8081")
}
