package router

import (
	"net/http"

	"github.com/diegodesousas/apistarter/di"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router    *httprouter.Router
	container di.Container
}

func New(configs ...ConfigRouter) Router {
	router := &Router{
		router: httprouter.New(),
	}

	for _, config := range configs {
		config(router)
	}

	return *router
}

func (r Router) AddRoute(route Route) {
	r.router.HandlerFunc(route.Method, route.Path, func(w http.ResponseWriter, req *http.Request) {
		route.Handler(w, req, r.container)
	})
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
