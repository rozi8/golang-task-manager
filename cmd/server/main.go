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

	taskHandler := handler.NewTaskHandler()
	mux.HandleFunc("/tasks", taskHandler.HandleTasks)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
