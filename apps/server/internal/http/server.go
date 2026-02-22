package http

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

func (s *Server) registerRoutes() {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
