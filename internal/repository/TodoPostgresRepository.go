package repository

import (
	"database/sql"
	"go-htmx-todo/internal/model"
	"time"
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

func (r *PostgreSQLTodoRepository) Create(todo model.Todo) (*model.Todo, error) {
	sqlStatement := `
		INSERT INTO todos (id, name, created_at) 
		VALUES ($1, $2, $3);
	`
	_, err := r.db.Exec(sqlStatement, todo.ID, todo.Name, todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *PostgreSQLTodoRepository) Complete(id string) error {
	sqlStatement := `
		UPDATE todo SET completed_at = $1 WHERE id = $2`
	_, err := r.db.Exec(sqlStatement, id, time.Now())

	if err != nil {
		return err
	}

	return nil
}
