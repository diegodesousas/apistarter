package middlewares

import (
	"log"
	"net/http"

	"github.com/diegodesousas/apistarter/di"
)

var GlobalMiddleware = func(container di.Container, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Global Middleware.")
		next.ServeHTTP(w, req)
	})
}
