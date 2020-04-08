package errorhandler

import (
	"net/http"

	"github.com/diegodesousas/apistarter/database"
)

type HTTPError struct {
	Err        error
	StatusCode int
}

func (e HTTPError) Error() string {
	return http.StatusText(e.StatusCode)
}

func NewHTTPError(err error, statusCode int) HTTPError {
	return HTTPError{Err: err, StatusCode: statusCode}
}

func httpStatusCode(err error) HTTPError {
	switch err {
	case database.NotFound:
		return NewHTTPError(err, http.StatusNotFound)
	}

	return NewHTTPError(err, http.StatusInternalServerError)
}

func HttpHandler(w http.ResponseWriter, err error) {
	httpError := httpStatusCode(err)
	http.Error(w, httpError.Error(), httpError.StatusCode)
}
