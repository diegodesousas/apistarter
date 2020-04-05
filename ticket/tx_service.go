package ticket

import (
	"log"

	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
)

type TxService interface {
	Create(*Ticket) error
}

type TxTicketService struct {
	tx           database.Transaction
	mediaService media.TxService
}

func NewTxTicketService(tx database.Transaction, txMediaService media.TxService) TxTicketService {
	return TxTicketService{
		tx:           tx,
		mediaService: txMediaService,
	}
}

func (t TxTicketService) Create(tkt *Ticket) error {
	if err := t.tx.Exec("INSERT INTO tickets ..."); err != nil {
		return err
	}

	log.Printf("created: %s", tkt)

	return t.mediaService.Create(tkt.ID, tkt.Medias)
}
