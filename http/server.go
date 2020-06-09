package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Server struct {
	server http.Server
}

func NewServer(configs ...ServerConfig) Server {
	s := &Server{
		server: http.Server{
			Addr: fmt.Sprintf(":%s", viper.GetString("HTTP_PORT")),
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
