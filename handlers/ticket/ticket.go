package ticket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/diegodesousas/apistarter/di"
	"github.com/diegodesousas/apistarter/ticket"
	"github.com/julienschmidt/httprouter"
)

var (
	FindByIdHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		FindById(w, r, container.NewTicketService())
	}
	CreateTicketHandler = func(w http.ResponseWriter, r *http.Request, container di.Container) {
		CreateTicket(w, r, container.NewTxlTicketService(container.NewTransaction()))
	}
)

func CreateTicket(w http.ResponseWriter, r *http.Request, service ticket.TxService) {
	tkt := &ticket.Ticket{}
	if err := json.NewDecoder(r.Body).Decode(tkt); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if err := service.Create(tkt); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func FindById(w http.ResponseWriter, r *http.Request, service ticket.Service) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	tkt, err := service.FindById(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(tkt); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
