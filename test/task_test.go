package test

import (
	"bytes"
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

func initMemoryDB(t *testing.T) (db *sql.DB, cleanup func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	// db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		db.Close()
	}
}

func TestListTasks(t *testing.T) {
	db, cleanup := initMemoryDB(t)
	defer cleanup()

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
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseTask := []task.Task{}
	err = json.NewDecoder(response.Body).Decode(&responseTask)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseTask)

	assert.Equal(t, []task.Task{task1}, responseTask)
	assert.Equal(t, task1.ID, responseTask[0].ID)
}

func TestCreateTask(t *testing.T) {
	db, cleanup := initMemoryDB(t)
	defer cleanup()

	migrate.Migrate(db)

	repository.TasksRepo = repository.NewTaskRepo(db)

	taskRequest := task.Task{
		ID:          uuid.New().String(),
		Title:       "some title",
		Description: "some desc",
		CompletedAt: time.Now().Format(time.RFC3339),
	}
	requestBody, err := json.Marshal(taskRequest)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(requestBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	mux := routes.NewTasksRouter()

	mux.ServeHTTP(rr, req)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	resJson := struct {
		Status int `json:"status"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&resJson)
	assert.NoError(t, err)

	taskReal := task.Task{}

	err = db.QueryRow("select id, title, description, completed_at from tasks where id = ? limit 1", taskRequest.ID).
		Scan(&taskReal.ID, &taskReal.Title, &taskReal.Description, &taskReal.CompletedAt)
	assert.NoError(t, err)

	assert.Equal(t, taskRequest.ID, taskReal.ID)
}
