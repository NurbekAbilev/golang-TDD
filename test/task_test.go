package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nurbekabilev/golang-tdd/internal/routes"
	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	mux := routes.NewTasksRouter()
	mux.ServeHTTP(rr, req)

	response := rr.Result()

	assert.Equal(t, response.StatusCode, http.StatusOK)
}
