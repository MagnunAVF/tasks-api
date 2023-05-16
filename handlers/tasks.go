package handlers

import (
	"net/http"

	"github.com/MagnunAVF/tasks-api/db"
	"github.com/MagnunAVF/tasks-api/models"
	"github.com/gin-gonic/gin"
)

func GetTasksHandler(c *gin.Context) {
	tasks, err := db.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
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
	// check attributes
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the task
	if err := db.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Task created successfully
	c.JSON(http.StatusOK, task)
}

func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
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
