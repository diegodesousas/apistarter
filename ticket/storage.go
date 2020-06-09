package ticket

import (
	"context"
)

type Storage interface {
	FindById(ctx context.Context, id string) (*Ticket, error)
	Create(ctx context.Context, tkt Ticket) error
}
