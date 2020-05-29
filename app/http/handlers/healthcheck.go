package handlers

import (
	"net/http"

	"github.com/diegodesousas/apistarter/core/di"
)

func Healthcheck(w http.ResponseWriter, _ *http.Request, _ di.Container) error {
	_, _ = w.Write([]byte("OK"))
	return nil
}
