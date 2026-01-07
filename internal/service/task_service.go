package service

import (
	"github.com/rozi8/golang-task-manager/internal/model"
	"github.com/rozi8/golang-task-manager/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() []model.Task {
	return s.repo.FindAll()
}

func (s *TaskService) CreateTask(title string) {
	task := model.Task{
		ID:     len(s.repo.FindAll()) + 1,
		Title:  title,
		Status: "pending",
	}
	s.repo.Save(task)
}
