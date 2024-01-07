package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.GetAllHandler)
	mux.HandleFunc("/outstanding", s.GetOutstandingHandler)
	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := s.todoRepository.GetAll()

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
