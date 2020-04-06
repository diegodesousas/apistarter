package media

import (
	"context"

	"github.com/diegodesousas/apistarter/database"
)

type Service interface {
	FindByTicketId(string) ([]Media, error)
}

type service struct {
	db database.Conn
}

func (s service) FindByTicketId(tid string) ([]Media, error) {
	var medias []Media

	sql := "SELECT * FROM medias WHERE ticket_id = $1"
	if err := s.db.SelectContext(context.Background(), &medias, sql, tid); err != nil {
		return medias, err
	}

	return medias, nil
}

func NewMediaService(db database.Conn) service {
	return service{
		db: db,
	}
}
