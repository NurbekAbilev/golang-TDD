package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nurbekabilev/golang-tdd/internal/app/repository"
	"github.com/nurbekabilev/golang-tdd/internal/data/task"
)

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	response := task.CreateTaskResponse{
		Status: http.StatusOK,
	}

	requestTask := task.Task{}
	err := json.NewDecoder(r.Body).Decode(&requestTask)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = repository.TasksRepo.CreateTask(r.Context(), requestTask)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(json)
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
