package server

import (
	"log"
	"net/http"
)

type Server struct {
	server http.Server
}

func New(configs ...ConfigServer) Server {
	s := &Server{
		server: http.Server{
			Addr: ":8080",
		},
	}

	for _, config := range configs {
		config(s)
	}

	return *s
}

func (s Server) Start() error {
	log.Printf("Starting server on %s", s.server.Addr)

	return s.server.ListenAndServe()
}
