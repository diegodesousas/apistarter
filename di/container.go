package di

import (
	"github.com/diegodesousas/apistarter/database"
	"github.com/diegodesousas/apistarter/media"
	"github.com/diegodesousas/apistarter/ticket"
)

type Container interface {
	NewTicketService() ticket.Service
	NewMediaService() media.Service
	NewTxMediaService(tx database.Transaction) media.TxService
	NewTxlTicketService(tx database.Transaction) ticket.TxService
	NewTransaction() database.Transaction
}

type defaultContainer struct{}

func (c defaultContainer) NewTxMediaService(tx database.Transaction) media.TxService {
	return media.NewTxService(tx)
}

func (c defaultContainer) NewTransaction() database.Transaction {
	return database.NewTx()
}

func (c defaultContainer) NewTxlTicketService(tx database.Transaction) ticket.TxService {
	return ticket.NewTxTicketService(tx, c.NewTxMediaService(tx))
}

func (c defaultContainer) NewMediaService() media.Service {
	return media.NewMediaService()
}

func (c defaultContainer) NewTicketService() ticket.Service {
	return ticket.NewService(c.NewMediaService())
}

func NewContainer() defaultContainer {
	return defaultContainer{}
}
