package test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nurbekabilev/golang-tdd/internal/app/repository"
	"github.com/nurbekabilev/golang-tdd/internal/data/task"
	"github.com/nurbekabilev/golang-tdd/internal/migrate"
	"github.com/nurbekabilev/golang-tdd/internal/routes"
	"github.com/stretchr/testify/assert"
)

func initMemoryDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	// db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestListTasks(t *testing.T) {
	db := initMemoryDB(t)

	migrate.Migrate(db)

	repository.TasksRepo = repository.NewTaskRepo(db)

	task1 := task.Task{
		ID:          uuid.New().String(),
		Title:       "some title",
		Description: "some desc",
		CompletedAt: time.Now().Format(time.RFC3339),
	}
	_, err := db.Exec("insert into tasks(id, title, description, completed_at) values (?,?,?,?)",
		task1.ID, task1.Title, task1.Description, task1.CompletedAt,
	)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	mux := routes.NewTasksRouter()
	mux.ServeHTTP(rr, req)

	response := rr.Result()

	assert.Equal(t, response.StatusCode, http.StatusOK)

	responseTask := []task.Task{}

	err = json.NewDecoder(response.Body).Decode(&responseTask)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseTask)

	assert.Equal(t, []task.Task{task1}, responseTask)
	assert.Equal(t, task1.ID, responseTask[0].ID)
}
