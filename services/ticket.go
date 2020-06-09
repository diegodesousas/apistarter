package services

import (
	"context"

	"github.com/diegodesousas/apistarter/ticket"
)

type TicketService struct {
	storage ticket.Storage
}

func NewTicketService(storage ticket.Storage) TicketService {
	return TicketService{
		storage: storage,
	}
}

func (ts TicketService) Create(ctx context.Context, ticket ticket.Ticket) error {
	//TODO validate ticket data
	return ts.storage.Create(ctx, ticket)
}

func (ts TicketService) FindById(ctx context.Context, id string) (*ticket.Ticket, error) {
	return ts.storage.FindById(ctx, id)
}
