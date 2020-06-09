package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/http/handlers"
)

var Routes = []Route{
	{
		Path:    "/healthcheck",
		Method:  http.MethodGet,
		Handler: handlers.Healthcheck,
	},
}
