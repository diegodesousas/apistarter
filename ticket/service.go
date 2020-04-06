package ticket

import (
	"context"

	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
)

type Service interface {
	FindById(string) (*Ticket, error)
}

func NewService(conn database.Conn, mediaService media.Service) DefaultTicketService {
	return DefaultTicketService{
		MediaService: mediaService,
		database:     conn,
	}
}

type DefaultTicketService struct {
	MediaService media.Service
	database     database.Conn
}

func (s DefaultTicketService) FindById(id string) (*Ticket, error) {
	tkt := &Ticket{}

	sql := "SELECT * FROM tickets WHERE id = $1"
	err := s.database.GetContext(context.Background(), tkt, sql, id)
	if err != nil {
		return nil, err
	}

	tkt.Medias, err = s.MediaService.FindByTicketId(tkt.ID)
	if err != nil {
		return nil, err
	}

	return tkt, nil
}
