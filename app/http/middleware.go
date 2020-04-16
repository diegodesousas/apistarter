package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/core/di"
	"github.com/justinas/alice"
)

var (
	Middlewares = func(middlewares ...Middleware) []Middleware {
		return append([]Middleware{}, middlewares...)
	}
)

type Middleware func(container di.Container, next http.Handler) http.Handler

func (m Middleware) Build(container di.Container) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return m(container, next)
	}
}
