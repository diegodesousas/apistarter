package handlers

import (
	"net/http"

	appHTTP "github.com/diegodesousas/apistarter/application/http"
)

var Routes = []appHTTP.Route{
	{
		Path:        "/tickets/:id",
		Method:      http.MethodGet,
		Handler:     FindTicketByIdHandler,
		Middlewares: appHTTP.Middlewares(ErrorMiddleware, RequestIDMiddleware),
	},
	{
		Path:    "/tickets",
		Method:  http.MethodPost,
		Handler: CreateTicketHandler,
	},
}
