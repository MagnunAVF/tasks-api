package db

import "github.com/MagnunAVF/tasks-api/models"

func CreateTask(task *models.Task) error {
	result := DB.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
