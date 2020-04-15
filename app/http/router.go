package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/app/database"
	"github.com/diegodesousas/apistarter/core/di"
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

func (r Router) addHandler(method string, path string, middlewares []Middleware, handler http.HandlerFunc) {
	r.router.Handler(
		method,
		path,
		alice.New(buildMiddlewares(r.container, middlewares...)...).ThenFunc(handler),
	)
}

func (r Router) AddRoute(route Route) {
	r.addHandler(
		route.Method,
		route.Path,
		route.Middlewares,
		func(w http.ResponseWriter, req *http.Request) {
			if err := route.Handler(w, req, r.container); err != nil {
				ErrorHandler(w, err)
			}
		})
}

func (r Router) AddTxRoute(route TxRoute) {
	r.addHandler(
		route.Method,
		route.Path,
		route.Middlewares,
		func(w http.ResponseWriter, req *http.Request) {
			conn, err := r.container.NewConn()
			if err != nil {
				ErrorHandler(w, err)
				return
			}

			err = conn.Transaction(func(tx database.TxConn) error {
				return route.Handler(w, req, tx, r.container)
			})

			if err != nil {
				ErrorHandler(w, err)
			}
		})
}

func buildMiddlewares(container di.Container, middlewares ...Middleware) []alice.Constructor {
	var list []alice.Constructor
	for _, middleware := range middlewares {
		list = append(list, middleware.Build(container))
	}
	return list
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.handler.ServeHTTP(w, req)
}
