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

func TestUpdateCard(t *testing.T) {
	type updateCard struct {
		mockBody model.Card
		mockID   int64
		mockOut  model.Card
		mockErr  error
	}
	type arg struct {
		updateCard           updateCard
		updateCardMockCalled bool
		givenBody            string
		givenID              string
		expRs                string
		expHTTPCode          int
	}
	tcs := map[string]arg{
		"success": {
			updateCard: updateCard{
				mockBody: model.Card{
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
				},
				mockID: 99,
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
			updateCardMockCalled: true,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
								}`,
			givenID: "99",
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
		"fail:error from service ": {
			updateCard: updateCard{
				mockBody: model.Card{
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
				},
				mockID:  99,
				mockOut: model.Card{},
				mockErr: errors.New("something wrong"),
			},
			updateCardMockCalled: true,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
							}`,
			givenID:     "99",
			expRs:       `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
		"fail: invalid number ": {
			updateCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203213215",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
							}`,
			givenID:     "99",
			expRs:       `{"code":"invalid_request", "description":"invalid number"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid expired date ": {
			updateCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2019-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
							}`,
			givenID:     "99",
			expRs:       `{"code":"invalid_request", "description":"invalid expired date"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid CVV": {
			updateCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "99999",
									"userID": 99,
									"balance": 9999
							}`,
			givenID:     "99",
			expRs:       `{"code":"invalid_request", "description":"invalid CVV"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid balance": {
			updateCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": -9999
							}`,
			givenID:     "99",
			expRs:       `{"code":"invalid_request", "description":"invalid balance"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid id": {
			updateCardMockCalled: false,
			givenBody: `{
									"number": "3301223454322203",
									"expired_date": "2023-03-22T00:00:00Z",
									"CVV": "999",
									"userID": 99,
									"balance": 9999
							}`,
			givenID:     "-99",
			expRs:       `{"code":"invalid_request", "description":"invalid id"}`,
			expHTTPCode: http.StatusBadRequest,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.updateCardMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("UpdateCard", mock.Anything, tc.updateCard.mockBody, tc.updateCard.mockID).
						Return(tc.updateCard.mockOut, tc.updateCard.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPut, "/cards", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.UpdateCard(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
