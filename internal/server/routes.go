package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"go-htmx-todo/internal/model"

	"github.com/google/uuid"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.GetAllHandler)
	//note to self
	//go uses the receiver pattern which registers these functions to the Server pointer
	mux.HandleFunc("/outstanding", s.GetOutstandingHandler)
	mux.HandleFunc("/add-todo", s.AddTodoHandler)
	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := s.todoRepository.GetAll()

	if err != nil {
		//todo - return appropriate error code to client
		log.Fatalf("Error fetching todos. Err: %v", err)
	}

	tmpl := template.Must(template.ParseFiles("./web/index.html"))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	tmpl.Execute(w, map[string][]model.Todo{"Todos": todos})
}

func (s *Server) AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/index.html"))
	tmpl.ExecuteTemplate(w, "todo-element", model.Todo{
		ID:        uuid.New(),
		Name:      "FooBar",
		CreatedAt: time.Now(),
	})
}

func (s *Server) GetOutstandingHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := s.todoRepository.GetOutstanding()

	if err != nil {
		//todo - return appropriate error code to client
		log.Fatalf("Error fetching todos. Err: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(todos)

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
