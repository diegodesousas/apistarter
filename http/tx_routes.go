package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/http/handlers"
)

var TxRoutes = []TxRoute{
	{
		Path:    "/tickets",
		Method:  http.MethodPost,
		Handler: handlers.CreateTicketHandler,
	},
}
