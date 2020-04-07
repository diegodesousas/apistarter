package ticket

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/ticket"
	"github.com/julienschmidt/httprouter"
)

var (
	findByIdHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		FindById(w, r, container.NewTicketService())
	}
	createTicketHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		conn, err := container.NewConn()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = conn.Transaction(func(tx database.TxConn) error {
			return CreateTicket(w, r, container.NewTxlTicketService(tx))
		})

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
)

func CreateTicket(w http.ResponseWriter, r *http.Request, service ticket.TxService) error {
	tkt := &ticket.Ticket{}
	if err := json.NewDecoder(r.Body).Decode(tkt); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	if err := service.Create(r.Context(), tkt); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	if err := json.NewEncoder(w).Encode(tkt); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	return nil
}

func FindById(w http.ResponseWriter, r *http.Request, service ticket.Service) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	tkt, err := service.FindById(id)
	if err == sql.ErrNoRows {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(tkt); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
