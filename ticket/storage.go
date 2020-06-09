package ticket

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
)

type Storage interface {
	FindById(ctx context.Context, id string) (*Ticket, error)
	FindMediaByTicketId(ctx context.Context, id string) ([]media.Media, error)
}

type storage struct {
	conn database.TxConn
}

func NewStorage(conn database.TxConn) storage {
	return storage{
		conn,
	}
}

func (s storage) FindById(ctx context.Context, id string) (*Ticket, error) {
	sql, args, err := squirrel.
		Select("*").
		From("tickets").
		Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	tkt := &Ticket{}
	err = s.conn.GetContext(ctx, tkt, sql, args...)
	if err != nil {
		return nil, database.HandleError(err)
	}

	return tkt, nil
}

func (s storage) FindMediaByTicketId(ctx context.Context, id string) ([]media.Media, error) {
	panic("implement me")
}
