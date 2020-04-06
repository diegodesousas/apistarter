package ticket

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
)

type TxService interface {
	Create(*Ticket) error
}

type TxTicketService struct {
	tx           database.TxConn
	mediaService media.TxService
}

func NewTxTicketService(tx database.TxConn, txMediaService media.TxService) TxTicketService {
	return TxTicketService{
		tx:           tx,
		mediaService: txMediaService,
	}
}

func (t TxTicketService) Create(tkt *Ticket) error {
	sql, args, err := squirrel.
		Insert("tickets").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		Columns("name").
		Values(tkt.Name).
		ToSql()

	if err = t.tx.ExecContext(context.Background(), sql, args...).Scan(&tkt.ID); err != nil {
		return err
	}

	for _, m := range tkt.Medias {
		if err := t.mediaService.Create(tkt.ID, &m); err != nil {
			return err
		}
	}

	return nil
}
