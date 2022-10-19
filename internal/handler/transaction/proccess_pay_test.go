package transaction

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	mocks "pg/internal/mocks/service/transaction"
	"pg/internal/model"
	"testing"
	"time"
)

func TestProcessPay(t *testing.T) {
	type processPay struct {
		mockTransID int64
		mockResp    model.Card
		mockErr     error
	}
	type arg struct {
		processPay           processPay
		processPayMockCalled bool
		givenID              string
		expRs                string
		expHTTPCode          int
	}
	tcs := map[string]arg{
		"success: ": {
			processPay: processPay{
				mockTransID: 100,
				mockResp: model.Card{
					ID:          100,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			processPayMockCalled: true,
			givenID:              "100",
			expRs:                "{\"balance\":9999, \"message\":\"Successful Transaction\", \"number\":\"3301223454322203\"}",
			expHTTPCode:          http.StatusOK,
		},
		"fail: invalid ID": {
			processPayMockCalled: false,
			givenID:              "-1000",
			expRs:                "{\"code\":\"invalid_request\", \"description\":\"invalid id\"}",
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail: error from service": {
			processPay: processPay{
				mockTransID: 100,
				mockResp: model.Card{
					ID:          100,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockErr: errors.New("something wrong"),
			},
			processPayMockCalled: true,
			givenID:              "100",
			expRs:                "{\"code\":\"internal_server_error\", \"description\":\"Something went wrong please try again later!\"}",
			expHTTPCode:          http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.processPayMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("InitPayment", mock.Anything, tc.processPay.mockTransID).
						Return(tc.processPay.mockResp, tc.processPay.mockErr),
				}
			}
			req := httptest.NewRequest(http.MethodPut, "/transaction/", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.ProcessPay(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
