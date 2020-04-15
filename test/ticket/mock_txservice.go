package ticket

import (
	"context"

	"github.com/diegodesousas/apistarter/core/ticket"
)

type MockTxService struct {
	MockCreate func(context.Context, *ticket.Ticket) error
}

func (m MockTxService) Create(ctx context.Context, tkt *ticket.Ticket) error {
	return m.MockCreate(ctx, tkt)
}
