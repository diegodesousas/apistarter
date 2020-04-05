package ticket

import (
	"github.com/diegodesousas/apistarter/media"
)

type Service interface {
	FindById(string) (Ticket, error)
}

func NewService(mediaService media.Service) DefaultTicketService {
	return DefaultTicketService{
		MediaService: mediaService,
	}
}

type DefaultTicketService struct {
	MediaService media.Service
}

func (s DefaultTicketService) FindById(id string) (Ticket, error) {
	return Ticket{
		ID:   id,
		Name: "Belford Roxo Tour 2020",
		Medias: []media.Media{
			{
				Path: "/media_1.jpg",
			},
			{
				Path: "/media_2.jpg",
			},
		},
	}, nil
}
