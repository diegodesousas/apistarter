package ticket

import "github.com/diegodesousas/apistarter/media"

type Ticket struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Medias []media.Media `json:"medias"`
}
