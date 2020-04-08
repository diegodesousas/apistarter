package ticket

import (
	"encoding/json"
	"net/http"

	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/errorhandler"
	"github.com/diegodesousas/apistarter/ticket"
	"github.com/julienschmidt/httprouter"
)

var (
	FindByIdHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		if err := FindById(w, r, container.NewTicketService()); err != nil {
			errorhandler.HttpHandler(w, err)
		}
	}
	CreateTicketHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		conn, err := container.NewConn()
		if err != nil {
			errorhandler.HttpHandler(w, err)
			return
		}

		err = conn.Transaction(func(tx database.TxConn) error {
			return CreateTicket(w, r, container.NewTxlTicketService(tx))
		})

		if err != nil {
			errorhandler.HttpHandler(w, err)
		}
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

func FindById(w http.ResponseWriter, req *http.Request, service ticket.Service) error {
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
