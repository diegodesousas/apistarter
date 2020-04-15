package handlers

import (
	"net/http"

	_ "github.com/diegodesousas/apistarter/application/http"
	infraHTTP "github.com/diegodesousas/apistarter/application/http"
)

var Routes = []infraHTTP.Route{
	{
		Path:        "/tickets/:id",
		Method:      http.MethodGet,
		Handler:     FindTicketByIdHandler,
		Middlewares: infraHTTP.Middlewares(ErrorMiddleware, RequestIDMiddleware),
	},
	{
		Path:    "/tickets",
		Method:  http.MethodPost,
		Handler: CreateTicketHandler,
	},
}
