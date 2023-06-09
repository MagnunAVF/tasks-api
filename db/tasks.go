package db

import (
	"github.com/MagnunAVF/tasks-api/models"
	"github.com/google/uuid"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByID(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := DB.First(&task, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(task *models.Task) error {
	// Generate UUID
	task.ID = uuid.New()
	result := DB.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateTaskById(id uuid.UUID, newTask *models.Task) (*models.Task, error) {
	var task models.Task
	if err := DB.First(&task, id).Error; err != nil {
		return nil, err
	}

	task.Title = newTask.Title
	task.Description = newTask.Description
	task.Done = newTask.Done
	task.Priority = newTask.Priority

	if err := DB.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteTaskByID(id uuid.UUID) error {
	if err := DB.Delete(&models.Task{}, "id = ?", id.String()).Error; err != nil {
		return err
	}

	return nil
}
