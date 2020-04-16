package middlewares

import (
	"log"
	"net/http"

	"github.com/diegodesousas/apistarter/core/di"
)

var ErrorMiddleware = func(container di.Container, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Error middleware.")
		next.ServeHTTP(w, req)
	})
}
