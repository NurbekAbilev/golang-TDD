package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nurbekabilev/golang-tdd/internal/data/task"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task task.Task) error
	GetTasks(ctx context.Context) ([]task.Task, error)
}

type TasksCreateRepository interface {
	CreateTask(ctx context.Context, task task.Task) error
}

type TasksGetRepository interface {
	GetTasks(ctx context.Context) ([]task.Task, error)
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
	_, err := r.db.Exec("insert into tasks (id, title, description, completed_at) values (?,?,?,?)",
		task.ID, task.Title, task.Description, task.CompletedAt)

	if err != nil {
		return fmt.Errorf("repo create task error: %w", err)
	}

	return nil
}

func (r *repo) GetTasks(ctx context.Context) ([]task.Task, error) {
	res := make([]task.Task, 0)

	rows, err := r.db.QueryContext(ctx, "select id, title, description, completed_at from tasks limit 10")
	if err != nil {
		return nil, fmt.Errorf("get tasks error: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CompletedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		res = append(res, t)
	}

	// todo log
	log.Printf("rows.next : %+v\n", res)

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return res, nil
}
