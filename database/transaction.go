package database

import (
	"log"
	"math/rand"
)

type Transaction interface {
	Exec(sql string) error
}

type Tx struct {
	id int
}

func (t Tx) Exec(sql string) error {
	log.Printf("tx id: %d persisted sql: %s", t.id, sql)
	return nil
}

func NewTx() *Tx {
	return &Tx{
		id: rand.Int(),
	}
}
