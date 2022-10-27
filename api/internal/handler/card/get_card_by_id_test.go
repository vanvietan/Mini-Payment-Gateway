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

func TestGetCardByID(t *testing.T) {
	type getCardByID struct {
		mockIn  int64
		mockOut model.Card
		mockErr error
	}
	type arg struct {
		getCardByID           getCardByID
		getCardByIDMockCalled bool
		givenID               string
		expRs                 string
		expHTTPCode           int
	}
	tcs := map[string]arg{
		"success": {
			getCardByID: getCardByID{
				mockIn: 99,
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
			getCardByIDMockCalled: true,
			givenID:               "99",
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
			getCardByID: getCardByID{
				mockIn:  99,
				mockOut: model.Card{},
				mockErr: errors.New("something wrong"),
			},
			getCardByIDMockCalled: true,
			givenID:               "99",
			expRs:                 `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode:           http.StatusInternalServerError,
		},
		"fail: invalid id": {
			getCardByIDMockCalled: false,
			givenID:               "-99",
			expRs:                 `{"code":"invalid_request", "description":"invalid id"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.getCardByIDMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetCardByID", mock.Anything, tc.getCardByID.mockIn).
						Return(tc.getCardByID.mockOut, tc.getCardByID.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodGet, "/cards", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetCardByID(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
