package main

import (
	"log"

	"github.com/diegodesousas/apistarter/application/http"
	"github.com/diegodesousas/apistarter/application/http/handlers"
	"github.com/diegodesousas/apistarter/core/di"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewRouter(
		http.WithContainer(container),
		http.WithRoutes(handlers.Routes...),
		http.WithTxRoutes(handlers.TxRoutes...),
		http.WithMiddleware(http.GlobalMiddleware),
	)

	s := http.NewServer(
		http.WithHandler(r),
	)

	log.Fatal(s.Start())
}
