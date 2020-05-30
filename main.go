package main

import (
	"log"

	"github.com/diegodesousas/apistarter/config"
	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/http"
	"github.com/diegodesousas/apistarter/http/middlewares"
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
		http.WithRoutes(http.Routes...),
		http.WithTxRoutes(http.TxRoutes...),
		http.WithMiddleware(middlewares.GlobalMiddleware),
	)

	s := http.NewServer(
		http.WithHandler(r),
	)

	log.Fatal(s.Start())
}
