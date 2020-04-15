package container

import (
	"github.com/diegodesousas/apistarter/application/database"
	"github.com/diegodesousas/apistarter/core/media"
	"github.com/diegodesousas/apistarter/core/ticket"
	testMedia "github.com/diegodesousas/apistarter/test/media"
	testTicket "github.com/diegodesousas/apistarter/test/ticket"
)

type MockContainer struct {
	MockTicketService   testTicket.MockService
	MockMediaService    testMedia.MockMediaService
	MockTxTicketService testTicket.MockTxService
	MockTxMediaService  testMedia.MockTxMediaService
	MockNewConn         func() (database.Conn, error)
}

func (m MockContainer) NewTicketService() ticket.Service {
	return m.MockTicketService
}

func (m MockContainer) NewMediaService(conn database.Conn) media.Service {
	panic("implement me")
}

func (m MockContainer) NewTxMediaService(tx database.TxConn) media.TxService {
	return m.MockTxMediaService
}

func (m MockContainer) NewTxlTicketService(tx database.TxConn) ticket.TxService {
	return m.MockTxTicketService
}

func (m MockContainer) NewConn() (database.Conn, error) {
	return m.MockNewConn()
}
