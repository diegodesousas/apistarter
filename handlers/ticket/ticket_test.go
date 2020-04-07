package ticket_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/diegodesousas/apistarter/handlers/ticket"
	"github.com/diegodesousas/apistarter/media"
	testTicket "github.com/diegodesousas/apistarter/test/ticket"
	"github.com/diegodesousas/apistarter/ticket"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicketShouldReturnSuccess(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "1"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	assert.Nil(t, err)

	expectedTicket := &ticket.Ticket{
		ID:   1,
		Name: "Ticket test",
		Medias: []media.Media{
			{
				ID:       1,
				Path:     "/test/media_1.jpg",
				TicketID: "1",
			},
		},
	}

	mockService := testTicket.MockService{
		FindByIdMocked: func(ctx context.Context, s string) (t *ticket.Ticket, err error) {
			return expectedTicket, nil
		},
	}

	response := httptest.NewRecorder()
	handler.FindById(response, request, mockService)
	assert.Equal(t, http.StatusOK, response.Code)

	bytes, err := json.Marshal(expectedTicket)
	assert.Nil(t, err)

	assert.Equal(t, string(bytes), response.Body.String())
}

func TestCreateTicketShouldReturnInternalServerError(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "0"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	assert.Nil(t, err)

	mockService := testTicket.MockService{
		FindByIdMocked: func(ctx context.Context, s string) (t *ticket.Ticket, err error) {
			return nil, errors.New("error")
		},
	}

	response := httptest.NewRecorder()
	handler.FindById(response, request, mockService)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError)+"\n", response.Body.String())
}

func TestCreateTicketShouldReturnNotFound(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "0"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	assert.Nil(t, err)

	mockService := testTicket.MockService{
		FindByIdMocked: func(ctx context.Context, s string) (t *ticket.Ticket, err error) {
			return nil, sql.ErrNoRows
		},
	}

	response := httptest.NewRecorder()
	handler.FindById(response, request, mockService)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, http.StatusText(http.StatusNotFound)+"\n", response.Body.String())
}
