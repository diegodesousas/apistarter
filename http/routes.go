package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/http/handlers"
	"github.com/diegodesousas/apistarter/http/middlewares"
)

var Routes = []Route{
	{
		Path:        "/tickets/:id",
		Method:      http.MethodGet,
		Handler:     handlers.FindTicketByIdHandler,
		Middlewares: Middlewares(middlewares.ErrorMiddleware, middlewares.RequestIDMiddleware),
	},
	{
		Path:    "/healthcheck",
		Method:  http.MethodGet,
		Handler: handlers.Healthcheck,
	},
}
