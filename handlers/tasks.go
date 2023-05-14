package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Priority    uint8  `json:"priority"`
}

func GetTasksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"route": "get tasks",
	})
}

func GetTaskHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"route": "get task",
		"id":    id,
	})
}

func CreateTaskHandler(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"route": "create task",
		"task":  task,
	})
}

func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"route": "create task",
		"task":  task,
		"id":    id,
	})
}

func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"route": "delete task",
		"id":    id,
	})
}
