package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Conn interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
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
