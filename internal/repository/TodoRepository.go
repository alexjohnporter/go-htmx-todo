package repository

import (
	"go-htmx-todo/internal/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetOutstanding() ([]model.Todo, error)
	// Create(todo model.Todo, error),
	// Complete(id uuid, error)
}
