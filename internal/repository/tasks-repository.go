package repository

import (
	"ai4am.com/go-restapi/internal/config"
	"ai4am.com/go-restapi/models"
	"strconv"
)

type TaskRepository struct{}
var taskRepository *TaskRepository

func GetTaskRepository() *TaskRepository {
	if taskRepository == nil {
		taskRepository = &TaskRepository{}
	}
	return taskRepository
}

func (r *TaskRepository) GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	where := models.Task{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &task, []string{"User", "Items"})
	if err != nil {
		return nil, err
	}
	return &task, err
}

func (r *TaskRepository) GetAllTasks() (*[]models.Task, error) {
	var tasks []models.Task
	err := Find(&models.Task{}, &tasks, []string{"User", "Items"}, "id asc")
	return &tasks, err
}

func (r *TaskRepository) QueryTasks(q *models.Task) (*[]models.Task, error) {
	var tasks []models.Task
	err := Find(&q, &tasks, []string{"User", "Items"}, "id asc")
	return &tasks, err
}

func (r *TaskRepository) AddTask(task *models.Task) error {
	err := Create(&task)
	err = Save(&task)
	return err
}

func (r *TaskRepository) UpdateTask(task *models.Task) error { return config.GetDB().Omit("User").Save(&task).Error }

func (r *TaskRepository) Delete(task *models.Task) error { return config.GetDB().Unscoped().Delete(&task).Error }
