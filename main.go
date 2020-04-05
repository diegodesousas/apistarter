package main

import (
	"log"

	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/handlers/ticket"
	"github.com/diegodesousas/apistarter/router"
	"github.com/diegodesousas/apistarter/server"
)

func main() {
	r := router.New(
		router.WithContainer(di.NewContainer()),
		router.WithRoutes(ticket.Routes()...),
	)

	s := server.New(
		server.WithHandler(r),
	)

	log.Fatal(s.Start())
}
