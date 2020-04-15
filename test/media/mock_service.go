package media

import (
	"context"

	"github.com/diegodesousas/apistarter/domain/media"
)

type MockMediaService struct {
	MockedFindByTicketId func(context.Context, int64) ([]media.Media, error)
}

func (m MockMediaService) FindByTicketId(ctx context.Context, tid int64) ([]media.Media, error) {
	return m.MockedFindByTicketId(ctx, tid)
}
