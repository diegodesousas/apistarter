package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type TxConn interface {
	Conn
	ExecContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Commit() error
	Rollback() error
}

type TxDatabase struct {
	Database
	*sqlx.Tx
}

func (db TxDatabase) ExecContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.QueryRowContext(ctx, query, args...)
}
