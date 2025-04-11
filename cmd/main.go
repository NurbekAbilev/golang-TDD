package main

import (
	"log"
	"os"

	"github.com/nurbekabilev/golang-tdd/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Printf("error occurred from app %v", err)
		os.Exit(1)
	}
}
