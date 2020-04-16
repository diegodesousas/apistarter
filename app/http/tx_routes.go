package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/app/http/handlers"
)

var TxRoutes = []TxRoute{
	{
		Path:    "/tickets",
		Method:  http.MethodPost,
		Handler: handlers.CreateTicketHandler,
	},
}
