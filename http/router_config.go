package http

import "github.com/diegodesousas/apistarter/di"

type RouterConfig func(*Router)

var (
	WithRoutes = func(routes ...Route) RouterConfig {
		return func(router *Router) {
			for _, route := range routes {
				router.AddRoute(route)
			}
		}
	}
	WithTxRoutes = func(routes ...TxRoute) RouterConfig {
		return func(router *Router) {
			for _, route := range routes {
				router.AddTxRoute(route)
			}
		}
	}
	WithContainer = func(container di.Container) RouterConfig {
		return func(router *Router) {
			router.container = container
		}
	}
	WithMiddleware = func(middlewares ...Middleware) RouterConfig {
		return func(router *Router) {
			router.middlewares = append(router.middlewares, middlewares...)
		}
	}
)
