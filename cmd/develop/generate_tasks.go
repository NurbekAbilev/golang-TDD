package develop

import (
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
			ID:          uuid.New(),
			Title:       "title 1",
			Description: "description 1",
			CompletedAt: time.Now(),
		},
		{
			ID:          uuid.New(),
			Title:       "title 2",
			Description: "description 2",
			CompletedAt: time.Time{},
		},
		{
			ID:          uuid.New(),
			Title:       "title 3",
			Description: "description 3",
			CompletedAt: time.Time{},
		},
	}

	for _, task := range tasks {
		_, err := db.Exec(
			"INSERT INTO tasks (id, title, description, completed_at) VALUES (?, ?, ?, ?)",
			task.ID, task.Title, task.Description, task.CompletedAt)
		if err != nil {
			panic(err)
		}
		// db.Exec("insert into tasks(id, title, description, values) values ($1, %2)")
	}
}
