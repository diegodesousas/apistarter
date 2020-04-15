package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/diegodesousas/apistarter/domain/media"
	"github.com/diegodesousas/apistarter/domain/ticket"
	"github.com/diegodesousas/apistarter/infra/database"
	"github.com/diegodesousas/apistarter/infra/http/handlers"
	"github.com/diegodesousas/apistarter/test/container"
	testDatabase "github.com/diegodesousas/apistarter/test/database"
	testMedia "github.com/diegodesousas/apistarter/test/media"
	testTicket "github.com/diegodesousas/apistarter/test/ticket"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestFindByIdShouldReturnSuccess(t *testing.T) {
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
	err = handlers.FindTicketById(response, request, mockService)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)

	bytes, err := json.Marshal(expectedTicket)
	assert.Nil(t, err)

	assert.Equal(t, string(bytes), response.Body.String())
}

func TestFindByIdShouldReturnInternalServerError(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "0"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	assert.Nil(t, err)

	mockTicketService := testTicket.MockService{
		FindByIdMocked: func(ctx context.Context, s string) (t *ticket.Ticket, err error) {
			return nil, errors.New("unknown error")
		},
	}
	mockMediaService := testMedia.MockMediaService{}

	mockContainer := container.MockContainer{
		MockTicketService: mockTicketService,
		MockMediaService:  mockMediaService,
	}

	response := httptest.NewRecorder()
	handlers.FindTicketByIdHandler(response, request, mockContainer)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError)+"\n", response.Body.String())
}

func TestFindByIdShouldReturnNotFound(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "0"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	assert.Nil(t, err)

	mockTicketService := testTicket.MockService{
		FindByIdMocked: func(ctx context.Context, s string) (t *ticket.Ticket, err error) {
			return nil, database.NotFound
		},
	}
	mockMediaService := testMedia.MockMediaService{}

	mockContainer := container.MockContainer{
		MockTicketService: mockTicketService,
		MockMediaService:  mockMediaService,
	}

	response := httptest.NewRecorder()
	handlers.FindTicketByIdHandler(response, request, mockContainer)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, http.StatusText(http.StatusNotFound)+"\n", response.Body.String())
}

func TestCreateTicketShouldReturnSuccess(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "1"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)

	expectedTicket := &ticket.Ticket{
		ID:   1,
		Name: "Ticket Test",
		Medias: []media.Media{
			{
				Path: "/test/media_1.jpg",
			},
			{
				Path: "/test/media_2.jpg",
			},
		},
	}

	mockTxService := testTicket.MockTxService{
		MockCreate: func(ctx context.Context, t *ticket.Ticket) error {
			t.ID = expectedTicket.ID
			return nil
		},
	}

	body := `
	{
		"name": "Ticket Test",
		"medias": [
			{
				"path": "/test/media_1.jpg"
			},
			{
				"path": "/test/media_2.jpg"
			}
		]
	}
	`
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/", strings.NewReader(body))
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	err = handlers.CreateTicket(response, request, mockTxService)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)

	bytes, err := json.Marshal(expectedTicket)
	assert.Nil(t, err)

	assert.Equal(t, string(bytes), response.Body.String())
}

func TestCreateTicketHandlerShouldReturnSuccess(t *testing.T) {
	params := httprouter.Params{{Key: "id", Value: "1"}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)

	expectedTicket := &ticket.Ticket{
		ID:   1,
		Name: "Ticket Test",
		Medias: []media.Media{
			{
				Path: "/test/media_1.jpg",
			},
			{
				Path: "/test/media_2.jpg",
			},
		},
	}

	mockTicketService := testTicket.MockTxService{
		MockCreate: func(ctx context.Context, t *ticket.Ticket) error {
			t.ID = expectedTicket.ID
			return nil
		},
	}
	mockTxMediaService := testMedia.MockTxMediaService{}

	mockContainer := container.MockContainer{
		MockTxTicketService: mockTicketService,
		MockTxMediaService:  mockTxMediaService,
		MockNewConn: func() (conn database.Conn, err error) {
			return testDatabase.SuccessTransaction, nil
		},
	}

	body := `
	{
		"name": "Ticket Test",
		"medias": [
			{
				"path": "/test/media_1.jpg"
			},
			{
				"path": "/test/media_2.jpg"
			}
		]
	}
	`
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/", strings.NewReader(body))
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	handlers.CreateTicketHandler(response, request, mockContainer)
	assert.Equal(t, http.StatusOK, response.Code)

	bytes, err := json.Marshal(expectedTicket)
	assert.Nil(t, err)

	assert.Equal(t, string(bytes), response.Body.String())
}
