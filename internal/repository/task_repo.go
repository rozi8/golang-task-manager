package repository

import "github.com/rozi8/golang-task-manager/internal/model"

type TaskRepository struct {
	tasks []model.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: []model.Task{},
	}
}

func (r *TaskRepository) FindAll() []model.Task {
	return r.tasks
}

func (r *TaskRepository) Save(task model.Task) {
	r.tasks = append(r.tasks, task)
}
