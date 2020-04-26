package api

import (
	"github.com/zbioe/ibcc/api/handlers"
	"net/http"
)

type server struct {
	mux *http.ServeMux
}

func NewServer() *server {
	return &server{
		mux: http.NewServeMux(),
	}
}

func (s *server) SetupHandlers() {
	s.mux.HandleFunc("/testMatrix", handlers.NewTestMatrix().Serve)
}

func (s *server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func (s *server) Run(port string) error {
	s.SetupHandlers()
	return s.Start(port)
}
