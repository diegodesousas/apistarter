package media

import (
	"log"

	"github.com/diegodesousas/apistarter/database"
)

type TxService interface {
	Create(tid string, medias []Media) error
}

type TxDefaultService struct {
	tx database.Transaction
}

func (t TxDefaultService) Create(tid string, medias []Media) error {
	log.Printf("created medias %s for ticket id %s", tid, medias)

	return t.tx.Exec("INSERT INTO medias...")
}

func NewTxService(tx database.Transaction) TxDefaultService {
	return TxDefaultService{
		tx: tx,
	}
}
