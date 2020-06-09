package ticket

import (
	"context"
)

type Service interface {
	Create(context.Context, Ticket) error
	FindById(context.Context, string) (*Ticket, error)
}
