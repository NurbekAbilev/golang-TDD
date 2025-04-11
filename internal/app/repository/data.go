package repository

import (
	"time"

	"github.com/google/uuid"
)

type Data struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CompletedAt time.Time `json:"completed"`
}
