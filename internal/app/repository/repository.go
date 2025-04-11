package repository

import (
	"context"
	"database/sql"

	"github.com/nurbekabilev/golang-tdd/internal/data/task"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task task.Task) error
	GetTask(ctx context.Context) ([]task.Task, error)
}

type TasksCreateRepository interface {
	CreateTask(ctx context.Context, task task.Task) error
}

type TasksGetRepository interface {
	GetTask(ctx context.Context) ([]task.Task, error)
}

var TasksRepo TaskRepository = nil

type repo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) TaskRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateTask(ctx context.Context, task task.Task) error {
	// r.db.Exec("insert into values")
	return nil
}

func (r *repo) GetTask(ctx context.Context) ([]task.Task, error) {
	panic("unimplemented")
}
