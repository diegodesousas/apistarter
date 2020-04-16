package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/app/http/handlers"
	"github.com/diegodesousas/apistarter/app/http/middlewares"
)

var Routes = []Route{
	{
		Path:        "/tickets/:id",
		Method:      http.MethodGet,
		Handler:     handlers.FindTicketByIdHandler,
		Middlewares: Middlewares(middlewares.ErrorMiddleware, middlewares.RequestIDMiddleware),
	},
}
