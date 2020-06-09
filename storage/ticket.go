package storage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
	"github.com/diegodesousas/apistarter/ticket"
)

type TicketStorage struct {
	conn database.TxConn
}

func (s TicketStorage) Create(ctx context.Context, tkt ticket.Ticket) error {
	sql, args, err := squirrel.Insert("tickets").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		Columns("name").
		Values(tkt.Name).
		ToSql()

	if err = s.conn.ExecContext(ctx, sql, args...).Scan(&tkt.ID); err != nil {
		return err
	}

	//for _, m := range tkt.Medias {
	//	if err := s.createMedia(ctx, tkt.ID, &m); err != nil {
	//		return err
	//	}
	//}

	return nil
}

func NewTicketStorage(conn database.TxConn) TicketStorage {
	return TicketStorage{
		conn,
	}
}

func (s TicketStorage) FindById(ctx context.Context, id string) (*ticket.Ticket, error) {
	sql, args, err := squirrel.
		Select("*").
		From("tickets").
		Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	tkt := &ticket.Ticket{}
	err = s.conn.GetContext(ctx, tkt, sql, args...)
	if err != nil {
		return nil, database.HandleError(err)
	}

	return tkt, nil
}

func (s TicketStorage) FindMediaByTicketId(ctx context.Context, id string) ([]media.Media, error) {
	panic("implement me")
}

func (s TicketStorage) createMedia(ctx context.Context, id int64, m *media.Media) error {
	panic("implement me")
}
