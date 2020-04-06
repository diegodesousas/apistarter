package di

import (
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
	"github.com/diegodesousas/apistarter/ticket"
)

type Container interface {
	NewTicketService() ticket.Service
	NewMediaService(conn database.Conn) media.Service
	NewTxMediaService(tx *database.Database) media.TxService
	NewTxlTicketService(tx *database.Database) ticket.TxService
	NewTransaction() database.Transaction
	NewConn() (database.Conn, error)
	Database() *database.Database
}

type container struct {
	conn database.Conn
}

func (c container) Database() *database.Database {
	return c.conn.(*database.Database)
}

func (c container) NewConn() (database.Conn, error) {
	return database.New("postgres", "postgres://postgres:root@postgres11.hud:5432/apistarter?sslmode=disable")
}

func (c container) NewTxMediaService(tx *database.Database) media.TxService {
	return media.NewTxService(tx)
}

func (c container) NewTransaction() database.Transaction {
	return database.NewTx()
}

func (c container) NewTxlTicketService(tx *database.Database) ticket.TxService {
	return ticket.NewTxTicketService(tx, c.NewTxMediaService(tx))
}

func (c container) NewMediaService(db database.Conn) media.Service {
	return media.NewMediaService(db)
}

func (c container) NewTicketService() ticket.Service {
	return ticket.NewService(c.conn, c.NewMediaService(c.conn))
}

func (c container) Build() (*container, error) {
	conn, err := c.NewConn()
	if err != nil {
		return nil, err
	}
	c.conn = conn

	return &c, nil
}

func NewContainer() (*container, error) {
	c := &container{}
	conn, err := c.NewConn()
	if err != nil {
		return nil, err
	}
	c.conn = conn

	return c, nil
}
