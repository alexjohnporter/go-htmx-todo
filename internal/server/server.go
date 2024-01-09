package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-htmx-todo/internal/database"
	"go-htmx-todo/internal/repository"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port           int
	db             database.Service
	todoRepository repository.TodoRepository
}

func NewServer(db database.Service, todoRepository repository.TodoRepository) (*http.Server, error) {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert PORT to integer: %v", err)
	}

	NewServer := &Server{
		port:           port,
		db:             db,
		todoRepository: todoRepository,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
