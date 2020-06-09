package ticket

import (
	"context"

	"github.com/diegodesousas/apistarter/media"
)

type Storage interface {
	FindById(ctx context.Context, id string) (*Ticket, error)
	FindMediaByTicketId(ctx context.Context, id string) ([]media.Media, error)
	Create(ctx context.Context, tkt Ticket) error
}
