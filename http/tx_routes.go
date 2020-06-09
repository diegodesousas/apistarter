package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/http/handlers"
	"github.com/diegodesousas/apistarter/http/middlewares"
)

var TxRoutes = []TxRoute{
	{
		Path:    "/tickets",
		Method:  http.MethodPost,
		Handler: handlers.CreateTicketHandler,
	},
	{
		Path:        "/tickets/:id",
		Method:      http.MethodGet,
		Handler:     handlers.FindTicketByIdHandler,
		Middlewares: Middlewares(middlewares.ErrorMiddleware, middlewares.RequestIDMiddleware),
	},
}
