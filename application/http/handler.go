package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/application/database"
	"github.com/diegodesousas/apistarter/core/di"
)

type Handler func(http.ResponseWriter, *http.Request, di.Container) error

type TxHandler func(http.ResponseWriter, *http.Request, database.TxConn, di.Container) error
