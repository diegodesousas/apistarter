package media

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/database"
)

type TxService interface {
	Create(context.Context, int64, *Media) error
}

type TxDefaultService struct {
	tx database.TxConn
}

func (t TxDefaultService) Create(ctx context.Context, tid int64, media *Media) error {
	sql, args, err := squirrel.
		Insert("medias").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		Columns("path", "ticket_id").
		Values(media.Path, tid).
		ToSql()

	if media.Path == "error" {
		return errors.New("test error")
	}
	if err != nil {
		return err
	}

	if err = t.tx.ExecContext(ctx, sql, args...).Scan(&media.ID); err != nil {
		return err
	}

	return nil
}

func NewTxService(tx database.TxConn) TxDefaultService {
	return TxDefaultService{
		tx: tx,
	}
}
