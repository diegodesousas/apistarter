package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/ticket"
	"github.com/julienschmidt/httprouter"
)

var (
	FindTicketByIdHandler = func(w http.ResponseWriter, r *http.Request, tx database.TxConn, container di.Container) error {
		return FindTicketById(w, r, container.NewTicketService(container.NewTicketStorage(tx)))
	}
	CreateTicketHandler = func(w http.ResponseWriter, r *http.Request, tx database.TxConn, container di.Container) error {
		return CreateTicket(w, r, container.NewTxlTicketService(tx))
	}
)

func CreateTicket(w http.ResponseWriter, r *http.Request, service ticket.TxService) error {
	tkt := &ticket.Ticket{}
	if err := json.NewDecoder(r.Body).Decode(tkt); err != nil {
		return err
	}

	if err := service.Create(r.Context(), tkt); err != nil {
		return err
	}

	bytes, err := json.Marshal(tkt)
	if err != nil {
		return err
	}
	if _, err := w.Write(bytes); err != nil {
		return err
	}

	return nil
}

func FindTicketById(w http.ResponseWriter, req *http.Request, service ticket.Service) error {
	ctx := req.Context()
	id := httprouter.ParamsFromContext(ctx).ByName("id")

	tkt, err := service.FindById(ctx, id)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(tkt)
	if err != nil {
		return err
	}

	if _, err = w.Write(bytes); err != nil {
		return err
	}

	return nil
}
