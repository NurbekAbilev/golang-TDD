package app

import (
	"log"
	"net/http"

	"github.com/nurbekabilev/golang-tdd/internal/app/conns"
	"github.com/nurbekabilev/golang-tdd/internal/app/repository"
	"github.com/nurbekabilev/golang-tdd/internal/migrate"
	"github.com/nurbekabilev/golang-tdd/internal/routes"
)

func Run() error {
	db, err := conns.InitSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}

	err = migrate.Migrate(db)
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
