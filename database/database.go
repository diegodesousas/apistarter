package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
