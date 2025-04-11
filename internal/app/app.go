package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/nurbekabilev/golang-tdd/internal/app/conns"
	"github.com/nurbekabilev/golang-tdd/internal/app/repository"
	"github.com/nurbekabilev/golang-tdd/internal/routes"
)

func Run() error {
	db, err := conns.InitSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}

	err = initdb(db)
	if err != nil {
		log.Fatal(err)
	}

	repository.TasksRepo = repository.NewTaskRepo(db)

	tasksRouter := routes.NewTasksRouter()

	http.Handle("/", tasksRouter)

	port := ":8080"
	log.Printf("listening http port %s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}

func initdb(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id UUID PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			completed_at TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
