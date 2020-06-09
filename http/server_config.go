package http

import "net/http"

type ServerConfig func(*Server)

var (
	WithHandler = func(handler http.Handler) ServerConfig {
		return func(s *Server) {
			s.server.Handler = handler
		}
	}
)
