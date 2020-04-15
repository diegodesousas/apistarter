package database

import (
	"context"
	"database/sql"

	"github.com/diegodesousas/apistarter/app/database"
)

type MockTxConn struct {
	MockConn
	MockExecContext func(ctx context.Context, query string, args ...interface{}) *sql.Row
	MockCommit      func() error
	MockRollback    func() error
}

func (m MockTxConn) ExecContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return m.MockExecContext(ctx, query, args...)
}

func (m MockTxConn) Commit() error {
	return m.MockCommit()
}

func (m MockTxConn) Rollback() error {
	return m.MockRollback()
}

var SuccessTransaction = MockConn{
	MockTransaction: func(f func(database.TxConn) error) error {
		tx := MockTxConn{
			MockCommit: func() error {
				return nil
			},
		}

		return f(tx)
	},
}
