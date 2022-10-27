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
	"strings"
	"testing"
	"time"
)

func TestAddCard(t *testing.T) {
	type addCard struct {
		mockIn  model.Card
		mockOut model.Card
		mockErr error
	}
	type arg struct {
		addCard           addCard
		addCardMockCalled bool
		givenBody         string
		expRs             string
		expHTTPCode       int
	}
	tcs := map[string]arg{
		"success: ": {
			addCard: addCard{
				mockIn: model.Card{
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
				},
				mockOut: model.Card{
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
			addCardMockCalled: true,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
								}`,
			expRs: `{				
									"ID" : 99,
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999,
									"created_at": "2022-03-14T14:00:00Z",
									"updated_at": "2022-03-14T14:00:00Z"
								}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: error from service": {
			addCard: addCard{
				mockIn: model.Card{
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
				},
				mockOut: model.Card{},
				mockErr: errors.New("something wrong"),
			},
			addCardMockCalled: true,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
								}`,
			expRs:       `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
		"fail: invalid number": {
			addCardMockCalled: false,
			givenBody: `{
									"number": "33012234543222035678899",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
								}`,
			expRs:       `{"code":"invalid_request", "description":"invalid number"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid expired date": {
			addCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2013-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
								}`,
			expRs:       `{"code":"invalid_request", "description":"invalid expired date"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid CVV": {
			addCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "99999",
									"userID": 99,
									"balance": 9999
								}`,
			expRs:       `{"code":"invalid_request", "description":"invalid CVV"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid balance": {
			addCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": -9999
								}`,
			expRs:       `{"code":"invalid_request", "description":"invalid balance"}`,
			expHTTPCode: http.StatusBadRequest,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.addCardMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("AddCard", mock.Anything, tc.addCard.mockIn).
						Return(tc.addCard.mockOut, tc.addCard.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/cards", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.AddCard(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
