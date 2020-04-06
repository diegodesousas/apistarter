package ticket

import (
	"net/http"

	"github.com/diegodesousas/apistarter/router"
)

func Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/tickets/:id",
			Method:  http.MethodGet,
			Handler: findByIdHandler,
		},
		{
			Path:    "/tickets",
			Method:  http.MethodPost,
			Handler: createTicketHandler,
		},
	}
}
