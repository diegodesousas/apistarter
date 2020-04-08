package database

import (
	"database/sql"
	"errors"
)

var (
	NotFound = errors.New("entity not found")
)

func HandleError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return NotFound

	default:
		return err
	}
}
