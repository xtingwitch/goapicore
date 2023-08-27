package coreHttp

import (
	"github.com/xtingwitch/GoApiCore/router"
	"net/http"
)

type Server struct {
	router *router.Router
	port   string
}

func NewServer(router *router.Router, port string) *Server {
	return &Server{
		router: router,
		port:   port,
	}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return server.ListenAndServe()
}
