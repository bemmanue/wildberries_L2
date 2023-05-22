package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func (s *Server) start(addr string) error {
	s.configureRouter()

	return http.ListenAndServe(addr, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/create_event", s.CreateEvent())
	s.router.HandleFunc("/update_event", s.UpdateEvent())
	s.router.HandleFunc("/delete_event", s.DeleteEvent())
	s.router.HandleFunc("/events_for_day", s.EventsForDay())
	s.router.HandleFunc("/events_for_week", s.EventsForWeek())
	s.router.HandleFunc("/events_for_month", s.EventsForMonth())
}

func (s *Server) CreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

func sendError(w http.ResponseWriter, code int, msg string) {
	doc := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doc)
}

func (s *Server) UpdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (s *Server) DeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (s *Server) EventsForDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (s *Server) EventsForWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (s *Server) EventsForMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func newServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

func main() {
	server := newServer()

	if err := server.start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
