package repository

import (
	"database/sql"
	"go-htmx-todo/internal/model"
)

type PostgreSQLTodoRepository struct {
	db *sql.DB
	//todo - what does this do?
	// mu      sync.RWMutex
}

func NewPostgresTodoRepository(db *sql.DB) *PostgreSQLTodoRepository {
	return &PostgreSQLTodoRepository{
		db: db,
	}
}

func (r *PostgreSQLTodoRepository) GetAll() ([]model.Todo, error) {

	rows, err := r.db.Query("SELECT id, name, created_at, completed_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *PostgreSQLTodoRepository) GetOutstanding() ([]model.Todo, error) {

	rows, err := r.db.Query("SELECT id, name, created_at, completed_at FROM todos WHERE completed_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// // GetByID retrieves a todo by its ID from the PostgreSQL repository
// func (r *PostgreSQLTodoRepository) GetByID(id int) (*model.Todo, error) {
// 	// Implementation...
// }

// // Create adds a new todo to the PostgreSQL repository
// func (r *PostgreSQLTodoRepository) Create(todo model.Todo) (*model.Todo, error) {
// 	// Implementation...
// }

// // Update modifies an existing todo in the PostgreSQL repository
// func (r *PostgreSQLTodoRepository) Update(id int, updatedTodo model.Todo) (*model.Todo, error) {
// 	// Implementation...
// }
