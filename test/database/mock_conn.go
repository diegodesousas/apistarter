package database

import (
	"context"

	"github.com/diegodesousas/apistarter/database"
)

type MockConn struct {
	MockedGetContext    func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MockedSelectContext func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MockedBegin         func() (MockTxConn, error)
	MockTransaction     func(func(database.TxConn) error) error
}

func (m MockConn) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.MockedGetContext(ctx, dest, query, args...)
}

func (m MockConn) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.MockedSelectContext(ctx, dest, query, args...)
}

func (m MockConn) Begin() (database.TxConn, error) {
	return m.MockedBegin()
}

func (m MockConn) Transaction(f func(database.TxConn) error) error {
	return m.MockTransaction(f)
}
