package card

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	mocks "pg/api/internal/mocks/service/card"
	"pg/api/internal/model"
	"testing"
	"time"
)

func TestGetCards(t *testing.T) {
	type getCards struct {
		mockLimit  int
		mockCursor int64
		mockOut    []model.Card
		mockErr    error
	}
	type arg struct {
		getCards           getCards
		getCardsMockCalled bool
		givenLimit         string
		givenCursor        string
		expRs              string
		expHTTPCode        int
	}
	tcs := map[string]arg{
		"success": {
			getCards: getCards{
				mockLimit:  20,
				mockCursor: 0,
				mockOut: []model.Card{
					{
						ID:          101,
						Number:      "3301223454322205",
						ExpiredDate: time.Date(2023, 3, 24, 00, 0, 0, 0, time.UTC),
						CVV:         "101",
						Balance:     10001,
						UserID:      101,
						CreatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
					},
					{
						ID:          100,
						Number:      "3301223454322204",
						ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
						CVV:         "100",
						Balance:     10000,
						UserID:      100,
						CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
					},
					{
						ID:          99,
						Number:      "3301223454322203",
						ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
						CVV:         "999",
						Balance:     9999,
						UserID:      99,
						CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					},
				},
			},
			getCardsMockCalled: true,
			givenLimit:         "20",
			givenCursor:        "0",
			expRs: `{
					"cards": [
						{
							"ID": 101,
							"number": "",
							"expired_date": "2023-03-24T00:00:00Z",
							"CVV": "101",
							"balance": 10001,
							"userID": 101,
							"created_at": "2022-03-16T14:00:00Z",
							"updated_at": "2022-03-16T14:00:00Z"
						},
						{
							"ID": 100,
							"number": "",
							"expired_date": "2023-03-23T00:00:00Z",
							"CVV": "100",
							"balance": 10000,
							"userID": 100,
							"created_at": "2022-03-15T14:00:00Z",
							"updated_at": "2022-03-15T14:00:00Z"
						},
						{
							"ID": 99,
							"number": "",
							"expired_date": "2023-03-22T00:00:00Z",
							"CVV": "999",
							"balance": 9999,
							"userID": 99,
							"created_at": "2022-03-14T14:00:00Z",
							"updated_at": "2022-03-14T14:00:00Z"
						}
					],
					"cursor": 99
				}`,
			expHTTPCode: http.StatusOK,
		},
		"success: empty result": {
			getCards: getCards{
				mockLimit:  20,
				mockCursor: 50,
				mockOut:    nil,
			},
			getCardsMockCalled: true,
			givenLimit:         "20",
			givenCursor:        "50",
			expRs:              `{"cards":null,"cursor":0}`,
			expHTTPCode:        http.StatusOK,
		},
		"fail: invalid limit": {
			getCardsMockCalled: false,
			givenLimit:         "200",
			givenCursor:        "50",
			expRs:              `{"code":"invalid_request", "description":"invalid limit"}`,
			expHTTPCode:        http.StatusBadRequest,
		},
		"fail: invalid cursor": {
			getCardsMockCalled: false,
			givenLimit:         "20",
			givenCursor:        "-50",
			expRs:              `{"code":"invalid_request", "description":"invalid cursor"}`,
			expHTTPCode:        http.StatusBadRequest,
		},
		"fail: error from service": {
			getCards: getCards{
				mockLimit:  20,
				mockCursor: 0,
				mockOut: []model.Card{
					{
						ID:          101,
						Number:      "3301223454322205",
						ExpiredDate: time.Date(2023, 3, 24, 00, 0, 0, 0, time.UTC),
						CVV:         "101",
						Balance:     10001,
						UserID:      101,
						CreatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
					},
					{
						ID:          100,
						Number:      "3301223454322204",
						ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
						CVV:         "100",
						Balance:     10000,
						UserID:      100,
						CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
					},
					{
						ID:          99,
						Number:      "3301223454322203",
						ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
						CVV:         "999",
						Balance:     9999,
						UserID:      99,
						CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					},
				},
				mockErr: errors.New("something wrong"),
			},
			getCardsMockCalled: true,
			givenLimit:         "20",
			givenCursor:        "0",
			expRs:              `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode:        http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.getCardsMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetCards", mock.Anything, tc.getCards.mockLimit, tc.getCards.mockCursor).
						Return(tc.getCards.mockOut, tc.getCards.mockErr),
				}
			}
			//GIVEN
			path := "/cards" + "?limit=" + tc.givenLimit + "&cursor=" + tc.givenCursor
			req := httptest.NewRequest(http.MethodGet, path, nil)
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetCards(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
