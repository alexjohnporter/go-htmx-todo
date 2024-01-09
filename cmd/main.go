package main

import (
	"fmt"
	"go-htmx-todo/internal/database"
	"go-htmx-todo/internal/repository"
	"go-htmx-todo/internal/server"
)

func main() {
	db := database.New()
	todoRepository := repository.NewPostgresTodoRepository(db.GetDB())
	server, _ := server.NewServer(db, todoRepository)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
