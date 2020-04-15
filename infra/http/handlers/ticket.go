package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diegodesousas/apistarter/domain/di"
	"github.com/diegodesousas/apistarter/domain/ticket"
	"github.com/diegodesousas/apistarter/infra/database"
	infraHTTP "github.com/diegodesousas/apistarter/infra/http"
	"github.com/julienschmidt/httprouter"
)

var (
	FindTicketByIdHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		if err := FindTicketById(w, r, container.NewTicketService()); err != nil {
			infraHTTP.ErrorHandler(w, err)
		}
	}
	CreateTicketHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		conn, err := container.NewConn()
		if err != nil {
			infraHTTP.ErrorHandler(w, err)
			return
		}

		err = conn.Transaction(func(tx database.TxConn) error {
			err := CreateTicket(w, r, container.NewTxlTicketService(tx))

			if err != nil {
				infraHTTP.ErrorHandler(w, err)
				return err
			}

			return nil
		})

		if err != nil {
			infraHTTP.ErrorHandler(w, err)
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
