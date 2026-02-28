package server

import (
	"net/http"
)

func (s *Server) registerRoutes() {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	s.router.Get("/helloWorld", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}