package app

import (
	"log"
	"net/http"
)

func Run() error {
	http.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye, World!"))
	})

	port := ":8080"
	log.Printf("listening http port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
}
