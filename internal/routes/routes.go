package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nurbekabilev/golang-tdd/internal/data/task"
)

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func handleListTasks(w http.ResponseWriter, r *http.Request) {
	taskID := uuid.New()
	resposne := []task.Task{
		{
			ID:          taskID,
			Title:       "Task title",
			Description: "Title description",
			CompletedAt: time.Time{},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resposne)
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
