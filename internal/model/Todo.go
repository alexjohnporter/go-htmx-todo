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

func (t Todo) FormatCreatedAt() string {
	return t.CreatedAt.Format(time.ANSIC)
}
