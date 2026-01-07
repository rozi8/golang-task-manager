package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rozi8/golang-task-manager/internal/repository"
	"github.com/rozi8/golang-task-manager/internal/service"
	"github.com/rozi8/golang-task-manager/pkg/response"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler() *TaskHandler {
	repo := repository.NewTaskRepository()
	svc := service.NewTaskService(repo)
	return &TaskHandler{service: svc}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := h.service.GetTasks()
		response.JSON(w, http.StatusOK, tasks)

	case http.MethodPost:
		var body struct {
			Title string `json:"title"`
		}

		json.NewDecoder(r.Body).Decode(&body)
		h.service.CreateTask(body.Title)
		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "task created",
		})
	}
}
