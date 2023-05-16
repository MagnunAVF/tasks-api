package main

import (
	"log"

	"github.com/MagnunAVF/tasks-api/db"
	"github.com/MagnunAVF/tasks-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", handlers.HealthHandler)

	router.GET("/tasks", handlers.GetTasksHandler)
	router.GET("/tasks/:id", handlers.GetTaskHandler)
	router.POST("/tasks", handlers.CreateTaskHandler)
	router.PUT("/tasks/:id", handlers.UpdateTaskHandler)
	router.DELETE("/tasks/:id", handlers.DeleteTaskHandler)

	router.Run(":8080")
}
