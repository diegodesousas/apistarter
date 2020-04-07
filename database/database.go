package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Conn interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Begin() (TxConn, error)
	Transaction(func(TxConn) error) error
}

type TxConn interface {
	Conn
	ExecContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Commit() error
	Rollback() error
}

type Database struct {
	*sqlx.DB
}

func New(driver string, conn string) (*Database, error) {
	db, err := sqlx.Open(driver, conn)
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func (db Database) Transaction(f func(TxConn) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = f(tx); err != nil {
		return tx.Commit()
	}

	return tx.Rollback()
}

type TxDatabase struct {
	Database
	*sqlx.Tx
}

func (db Database) Begin() (TxConn, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	return &TxDatabase{db, tx}, nil
}

func (db TxDatabase) ExecContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.QueryRowContext(ctx, query, args...)
}
