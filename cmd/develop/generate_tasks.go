package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nurbekabilev/golang-tdd/internal/app/conns"
	"github.com/nurbekabilev/golang-tdd/internal/data/task"
)

func main() {
	db, err := conns.InitSQLiteConn()
	if err != nil {
		panic(err)
	}

	tasks := []task.Task{
		{
			ID:          uuid.New().String(),
			Title:       "title 1",
			Description: "description 1",
			CompletedAt: time.Now().Format(time.RFC3339),
		},
		{
			ID:          uuid.New().String(),
			Title:       "title 2",
			Description: "description 2",
			CompletedAt: "",
		},
		{
			ID:          uuid.New().String(),
			Title:       "title 3",
			Description: "description 3",
			CompletedAt: "",
		},
	}

	for _, task := range tasks {
		log.Printf("running seed for %s\n", task.Title)
		_, err := db.Exec(
			"INSERT INTO tasks (id, title, description, completed_at) VALUES (?, ?, ?, ?)",
			task.ID, task.Title, task.Description, task.CompletedAt)
		if err != nil {
			log.Printf("error occured for task %+v %w\n", task, err)
		}
	}
}
