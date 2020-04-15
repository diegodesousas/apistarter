package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/domain/di"
)

type Handler func(http.ResponseWriter, *http.Request, di.Container)
