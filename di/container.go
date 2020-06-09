package di

import (
	"github.com/diegodesousas/apistarter/config"
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
	"github.com/diegodesousas/apistarter/services"
	"github.com/diegodesousas/apistarter/storage"
	"github.com/diegodesousas/apistarter/ticket"
)

type Container interface {
	NewTicketService(ticket.Storage) ticket.Service
	NewMediaService(database.Conn) media.Service
	NewTxMediaService(database.TxConn) media.TxService
	NewTicketStorage(database.TxConn) ticket.Storage
	NewConn() (database.Conn, error)
	NewTxConn() (database.TxConn, error)
}

type container struct {
	conn   database.Conn
	config *config.Config
}

func (c container) NewTicketService(storage ticket.Storage) ticket.Service {
	return services.NewTicketService(storage)
}

func (c container) NewTicketStorage(tx database.TxConn) ticket.Storage {
	return storage.NewTicketStorage(tx)
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

func (c container) NewMediaService(db database.Conn) media.Service {
	return media.NewMediaService(db)
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
