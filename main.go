package main

import (
	"github.com/MagnunAVF/tasks-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handlers.HealthHandler)

	router.Run(":8080")
}
