package router

import "github.com/diegodesousas/apistarter/di"

type ConfigRouter func(*Router)

var (
	WithRoutes = func(routes ...Route) ConfigRouter {
		return func(router *Router) {
			for _, route := range routes {
				router.AddRoute(route)
			}
		}
	}
	WithContainer = func(container di.Container) ConfigRouter {
		return func(router *Router) {
			router.container = container
		}
	}
)
