package handlers

import (
	"net/http"

	appHttp "github.com/diegodesousas/apistarter/application/http"
)

var TxRoutes = []appHttp.TxRoute{
	{
		Path:    "/tickets",
		Method:  http.MethodPost,
		Handler: CreateTicketHandler,
	},
}
