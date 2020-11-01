package http

import (
	"github.com/JamieBShaw/auth-service/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Server interface {
	Create(rw http.ResponseWriter, r *http.Request)
}

type httpServer struct {
	service service.AuthService
	router *mux.Router
}

func (s *httpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(rw, r)
}

func NewHttpServer(service service.AuthService, router *mux.Router) http.Handler {
	server := &httpServer{service: service, router: router}
	server.routes()

	return server
}


func (s *httpServer) Create(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented")
}


