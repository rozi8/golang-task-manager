package main

import (
	"log"
	"net/http"

	"github.com/rozi8/golang-task-manager/internal/config"
	"github.com/rozi8/golang-task-manager/internal/handler"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task Manager API berjalan OK"))
	})

	taskHandler := handler.NewTaskHandler()
	mux.HandleFunc("/tasks", taskHandler.HandleTasks)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
