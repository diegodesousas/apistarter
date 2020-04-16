package media

import (
	"context"

	"github.com/diegodesousas/apistarter/core/media"
)

type MockTxMediaService struct {
	MockCreate func(ctx context.Context, tid int64, media *media.Media) error
}

func (m MockTxMediaService) Create(ctx context.Context, tid int64, media *media.Media) error {
	return m.MockCreate(ctx, tid, media)
}
