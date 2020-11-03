package http

import (
	"encoding/json"
	"errors"
	"github.com/JamieBShaw/auth-service/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type Server interface {
	Create(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
}

type httpServer struct {
	service service.AuthService
	router  *mux.Router
	log     *logrus.Logger
}

func (s *httpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(rw, r)
}

func NewHttpServer(service service.AuthService, router *mux.Router, log *logrus.Logger) http.Handler {
	server := &httpServer{service: service, router: router, log: log}
	server.routes()

	return server
}

func (s *httpServer) Create(rw http.ResponseWriter, r *http.Request) {
	s.log.Info("[HTTP SERVER] Executing Create Handler")

	userId := strings.TrimSpace(mux.Vars(r)["id"])

	id, err := strconv.Atoi(userId)
	if err != nil {
		s.log.Errorf("error: %v", err.Error())
		http.Error(rw, errors.New("invalid request").Error(), http.StatusBadRequest)
		return
	}

	authTokens, err := s.service.Create(int64(id))
	if err != nil {
		s.log.Errorf("error: %v", err.Error())
		http.Error(rw, errors.New("could not create access token for user").Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(rw).Encode(&authTokens)
	if err != nil {
		s.log.Errorf("error: %v", err.Error())
		http.Error(rw, errors.New("could not create access token for user").Error(), http.StatusInternalServerError)
	}
}

func (s *httpServer) Delete(rw http.ResponseWriter, r *http.Request) {
	panic("not yet implemented")
}


