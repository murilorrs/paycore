package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	r := chi.NewRouter()

	s := &Server{
		router: r,
	}

	s.registerRoutes()

	return s
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
