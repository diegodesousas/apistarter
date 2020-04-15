package media

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/infra/database"
)

type TxService interface {
	Create(context.Context, int64, *Media) error
}

type txService struct {
	tx database.TxConn
}

func (t txService) Create(ctx context.Context, tid int64, media *Media) error {
	sql, args, err := squirrel.
		Insert("medias").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		Columns("path", "ticket_id").
		Values(media.Path, tid).
		ToSql()

	if media.Path == "errorhandler" {
		return errors.New("test errorhandler")
	}
	if err != nil {
		return err
	}

	if err = t.tx.ExecContext(ctx, sql, args...).Scan(&media.ID); err != nil {
		return err
	}

	return nil
}

func NewTxService(tx database.TxConn) txService {
	return txService{
		tx: tx,
	}
}
