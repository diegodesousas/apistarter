package main

import (
	"log"

	"github.com/diegodesousas/apistarter/app/config"
	"github.com/diegodesousas/apistarter/app/http"
	"github.com/diegodesousas/apistarter/app/http/handlers"
	"github.com/diegodesousas/apistarter/core/di"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	container, err := di.NewContainer(cfg)
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
