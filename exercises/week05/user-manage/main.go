package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user_manage/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users", handlers.UpdateUser)
	router.DELETE("/users/:email", handlers.DeleteUser)
	log.Println("服务器启动，监听端口 :8080")
	router.Run(":8080")
}
