package main

import (
	"log"
	"os"

	"github.com/MagnunAVF/tasks-api/db"
	"github.com/MagnunAVF/tasks-api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

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
