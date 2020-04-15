package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/domain/di"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type Router struct {
	router      *httprouter.Router
	container   di.Container
	middlewares []Middleware
	handler     http.Handler
}

func NewRouter(configs ...RouterConfig) Router {
	router := &Router{
		router: httprouter.New(),
	}

	for _, config := range configs {
		config(router)
	}

	router.handler = alice.
		New(buildMiddlewares(router.container, router.middlewares...)...).
		Then(router.router)

	return *router
}

func (r Router) AddRoute(route Route) {
	main := func(w http.ResponseWriter, req *http.Request) {
		route.Handler(w, req, r.container)
	}

	middlewares := buildMiddlewares(r.container, route.Middlewares...)

	r.router.Handler(
		route.Method,
		route.Path,
		alice.New(middlewares...).ThenFunc(main),
	)
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.handler.ServeHTTP(w, req)
}

func buildMiddlewares(container di.Container, middlewares ...Middleware) []alice.Constructor {
	var list []alice.Constructor
	for _, middleware := range middlewares {
		list = append(list, middleware.Build(container))
	}
	return list
}
