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

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()
	NewServer := &Server{
		port:           port,
		db:             db,
		todoRepository: repository.NewPostgresTodoRepository(db.GetDB()),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
