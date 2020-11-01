package http

import "net/http"

func (s *httpServer) routes() {

	post := s.router.Methods(http.MethodPost).Subrouter()

	post.HandleFunc("/auth/create/{id}", s.Create)
}