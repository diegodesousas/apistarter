package handlers

import (
	"log"
	"net/http"

	"github.com/diegodesousas/apistarter/core/di"
)

var (
	RequestIDMiddleware = func(container di.Container, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Println("Generate request id and put in context.")
			next.ServeHTTP(w, req)
		})
	}
	ErrorMiddleware = func(container di.Container, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Println("Error middleware.")
			next.ServeHTTP(w, req)
		})
	}
)
