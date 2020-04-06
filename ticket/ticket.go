package ticket

import "github.com/diegodesousas/apistarter/media"

type Ticket struct {
	ID     int64         `json:"id" db:"id"`
	Name   string        `json:"name" db:"name"`
	Medias []media.Media `json:"medias" db:"-"`
}
