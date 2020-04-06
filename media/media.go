package media

type Media struct {
	ID       int64  `json:"-" db:"id"`
	Path     string `json:"path" db:"path"`
	TicketID string `json:"-" db:"ticket_id"`
}
