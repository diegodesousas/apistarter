package http

import (
	"net/http"

	"github.com/diegodesousas/apistarter/app/database"
)

type Error struct {
	Err        error
	StatusCode int
}

func (e Error) Error() string {
	return http.StatusText(e.StatusCode)
}

func NewHTTPError(err error, statusCode int) Error {
	return Error{Err: err, StatusCode: statusCode}
}

func httpStatusCode(err error) Error {
	switch err.(type) {
	case database.NotFoundErr:
		return NewHTTPError(err, http.StatusNotFound)
	default:
		return NewHTTPError(err, http.StatusInternalServerError)
	}
}

func ErrorHandler(w http.ResponseWriter, err error) {
	httpError := httpStatusCode(err)
	http.Error(w, httpError.Error(), httpError.StatusCode)
}
