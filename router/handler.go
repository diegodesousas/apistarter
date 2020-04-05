package router

import (
	"net/http"

	"github.com/diegodesousas/apistarter/di"
)

type Handler func(http.ResponseWriter, *http.Request, di.Container)
