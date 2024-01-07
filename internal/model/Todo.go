package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	Name        string
	CreatedAt   time.Time
	CompletedAt *time.Time
}
