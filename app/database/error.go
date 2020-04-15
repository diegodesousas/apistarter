package database

import (
	"database/sql"
)

func HandleError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return NewNotFoundErr()

	default:
		return err
	}
}
