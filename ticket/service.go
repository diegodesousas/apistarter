package ticket

import (
	"context"
)

type Service interface {
	FindById(context.Context, string) (*Ticket, error)
}

func NewService(storage Storage) service {
	return service{
		storage,
	}
}

type service struct {
	storage Storage
}

func (s service) FindById(ctx context.Context, id string) (*Ticket, error) {
	return s.storage.FindById(ctx, id)
}
