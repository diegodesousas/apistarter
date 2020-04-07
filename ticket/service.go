package ticket

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
)

type Service interface {
	FindById(context.Context, string) (*Ticket, error)
}

func NewService(conn database.Conn, mediaService media.Service) service {
	return service{
		mediaService: mediaService,
		database:     conn,
	}
}

type service struct {
	mediaService media.Service
	database     database.Conn
}

func (s service) FindById(ctx context.Context, id string) (*Ticket, error) {
	tkt := &Ticket{}

	sql, args, err := squirrel.
		Select("*").
		From("tickets").
		Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = s.database.GetContext(ctx, tkt, sql, args...)
	if err != nil {
		return nil, err
	}

	tkt.Medias, err = s.mediaService.FindByTicketId(ctx, tkt.ID)
	if err != nil {
		return nil, err
	}

	return tkt, nil
}
