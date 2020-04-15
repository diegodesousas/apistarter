package ticket

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/application/database"
	"github.com/diegodesousas/apistarter/core/media"
)

type TxService interface {
	Create(context.Context, *Ticket) error
}

type txService struct {
	tx           database.TxConn
	mediaService media.TxService
}

func NewTxTicketService(tx database.TxConn, txMediaService media.TxService) txService {
	return txService{
		tx:           tx,
		mediaService: txMediaService,
	}
}

func (t txService) Create(ctx context.Context, tkt *Ticket) error {
	sql, args, err := squirrel.
		Insert("tickets").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		Columns("name").
		Values(tkt.Name).
		ToSql()

	if err = t.tx.ExecContext(ctx, sql, args...).Scan(&tkt.ID); err != nil {
		return err
	}

	for _, m := range tkt.Medias {
		if err := t.mediaService.Create(ctx, tkt.ID, &m); err != nil {
			return err
		}
	}

	return nil
}
