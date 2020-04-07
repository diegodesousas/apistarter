package router

import (
	"log"
	"net/http"

	"github.com/diegodesousas/apistarter/di"
	"github.com/justinas/alice"
)

var (
	Middlewares = func(middlewares ...Middleware) []Middleware {
		return append([]Middleware{}, middlewares...)
	}
	GlobalMiddleware = func(container di.Container, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Println("Global Middleware.")
			next.ServeHTTP(w, req)
		})
	}
)

type Middleware func(container di.Container, next http.Handler) http.Handler

func (m Middleware) Build(container di.Container) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return m(container, next)
	}
}
