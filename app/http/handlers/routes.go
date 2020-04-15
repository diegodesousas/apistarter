package handlers

import (
	"net/http"

	appHTTP "github.com/diegodesousas/apistarter/app/http"
)

var Routes = []appHTTP.Route{
	{
		Path:        "/tickets/:id",
		Method:      http.MethodGet,
		Handler:     FindTicketByIdHandler,
		Middlewares: appHTTP.Middlewares(ErrorMiddleware, RequestIDMiddleware),
	},
}
