package data

import (
	"errors"
	"fmt"
	"sync"
	"task_manager/models"
	"time"
)

type TaskService struct {
	mu     sync.RWMutex
	tasks  map[string]models.Task
	nextID int
}

func NewTaskService() *TaskService {
	service := &TaskService{
		tasks:  make(map[string]models.Task),
		nextID: 1,
	}
	// Seed initial data
	service.CreateTask(models.Task{
		Title:       "Set up Go Gin REST API",
		Description: "Build a Task Management API using Gin framework",
		DueDate:     time.Now().AddDate(0, 0, 7),
		Status:      "In Progress",
	})
	return service
}

func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	taskList := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = fmt.Sprintf("%d", s.nextID)
	s.nextID++

	if task.Status == "" {
		task.Status = "Pending"
	}

	s.tasks[task.ID] = task
	return task
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	updatedTask.ID = id
	s.tasks[id] = updatedTask
	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}