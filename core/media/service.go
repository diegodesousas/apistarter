package media

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/diegodesousas/apistarter/app/database"
)

type Service interface {
	FindByTicketId(context.Context, int64) ([]Media, error)
}

type service struct {
	db database.Conn
}

func (s service) FindByTicketId(ctx context.Context, tid int64) ([]Media, error) {
	var medias []Media

	sql, args, err := squirrel.
		Select("*").
		From("medias").
		Where("ticket_id = ?", tid).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err = s.db.SelectContext(ctx, &medias, sql, args...); err != nil {
		return medias, err
	}

	return medias, nil
}

func NewMediaService(db database.Conn) service {
	return service{
		db: db,
	}
}
