package media

import (
	"context"

	"github.com/diegodesousas/apistarter/database"
)

type Service interface {
	FindByTicketId(string) ([]Media, error)
}

type service struct {
	db *database.Database
}

func (s service) FindByTicketId(tid string) ([]Media, error) {
	medias := []Media{}

	sql := "SELECT * FROM medias WHERE ticket_id = $1"
	if err := s.db.SelectContext(context.Background(), &medias, sql, tid); err != nil {
		return []Media{}, err
	}

	return medias, nil
}

func NewMediaService(db *database.Database) service {
	return service{
		db: db,
	}
}
