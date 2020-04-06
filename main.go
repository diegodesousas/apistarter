package main

import (
	"log"

	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/handlers/ticket"
	"github.com/diegodesousas/apistarter/router"
	"github.com/diegodesousas/apistarter/server"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	r := router.New(
		router.WithContainer(container),
		router.WithRoutes(ticket.Routes()...),
	)

	s := server.New(
		server.WithHandler(r),
	)

	log.Fatal(s.Start())
}
