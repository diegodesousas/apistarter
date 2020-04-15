package main

import (
	"log"

	"github.com/diegodesousas/apistarter/domain/di"
	"github.com/diegodesousas/apistarter/infra/http"
	"github.com/diegodesousas/apistarter/infra/http/handlers"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewRouter(
		http.WithContainer(container),
		http.WithRoutes(handlers.Routes...),
		http.WithMiddleware(http.GlobalMiddleware),
	)

	s := http.NewServer(
		http.WithHandler(r),
	)

	log.Fatal(s.Start())
}
