package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nurbekabilev/golang-tdd/internal/app/repository"
)

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func handleListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.TasksRepo.GetTasks(r.Context())
	if err != nil {
		log.Printf("error occured in handleListTasks: %s", err)
		http.Error(w, "error occured", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}

func NewTasksRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handleListTasks)
	mux.HandleFunc("POST /tasks", handleCreateTask)
	return mux
}
