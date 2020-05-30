package di

import (
	"github.com/diegodesousas/apistarter/config"
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
	"github.com/diegodesousas/apistarter/ticket"
)

type Container interface {
	NewTicketService() ticket.Service
	NewMediaService(conn database.Conn) media.Service
	NewTxMediaService(tx database.TxConn) media.TxService
	NewTxlTicketService(tx database.TxConn) ticket.TxService
	NewConn() (database.Conn, error)
}

type container struct {
	conn   database.Conn
	config *config.Config
}

func (c container) NewTxConn() (database.TxConn, error) {
	return c.conn.Begin()
}

func (c container) NewConn() (database.Conn, error) {
	return database.New(c.config.Database.Driver, c.config.Database.DSN)
}

func (c container) NewTxMediaService(tx database.TxConn) media.TxService {
	return media.NewTxService(tx)
}

func (c container) NewTxlTicketService(tx database.TxConn) ticket.TxService {
	return ticket.NewTxTicketService(tx, c.NewTxMediaService(tx))
}

func (c container) NewMediaService(db database.Conn) media.Service {
	return media.NewMediaService(db)
}

func (c container) NewTicketService() ticket.Service {
	return ticket.NewService(c.conn, c.NewMediaService(c.conn))
}

func NewContainer(cfg *config.Config) (*container, error) {
	c := &container{
		config: cfg,
	}
	conn, err := c.NewConn()
	if err != nil {
		return nil, err
	}
	c.conn = conn

	return c, nil
}
