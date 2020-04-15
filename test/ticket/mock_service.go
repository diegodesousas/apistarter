package ticket

import (
	"context"

	"github.com/diegodesousas/apistarter/core/ticket"
)

type MockService struct {
	FindByIdMocked func(context.Context, string) (*ticket.Ticket, error)
}

func (m MockService) FindById(ctx context.Context, id string) (*ticket.Ticket, error) {
	return m.FindByIdMocked(ctx, id)
}
