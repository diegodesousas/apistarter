package server

import "net/http"

type ConfigServer func(*Server)

var (
	WithHandler = func(handler http.Handler) ConfigServer {
		return func(s *Server) {
			s.server.Handler = handler
		}
	}
)
