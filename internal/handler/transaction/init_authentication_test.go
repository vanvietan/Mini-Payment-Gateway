package transaction

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	mocks "pg/internal/mocks/service/transaction"
	"pg/internal/model"
	"strings"
	"testing"
	"time"
)

func TestInitAuthentication(t *testing.T) {
	type initAuthentication struct {
		mockInCard   model.Card
		mockInOrder  model.Order
		mockOutCard  model.Card
		mockOutOrder model.Order
		mockOutTrans model.Transaction
		mockErr      error
		mockErrB     error
	}
	type arg struct {
		initAuthentication           initAuthentication
		initAuthenticationMockCalled bool
		givenBody                    string
		expRs                        string
		expHTTPCode                  int
	}
	tcs := map[string]arg{
		"success: ": {
			initAuthentication: initAuthentication{
				mockInCard: model.Card{
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
				},
				mockInOrder: model.Order{
					Amount: 99,
				},
				mockOutCard: model.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOutOrder: model.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOutTrans: model.Transaction{
					ID:        99,
					CardID:    99,
					OrderID:   99,
					OTP:       "123456",
					Status:    model.StatusPending,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			initAuthenticationMockCalled: true,
			givenBody: `{
							"number": "3301223454322203",
							"expired_date": "2023-03-22T00:00:00Z",
							"CVV": "999",
							"amount" : 99
						} `,
			expRs: `{
				"message": "created a transaction",
				"html": "<!DOCTYPE html>\n<html>\n<body>\n<h1>Submit your OTP</h1>\n<form action=\"/transactions\" method=\"post\">\n    <label for=\"otp\">OTP:</label>\n    <input type=\"text\" id=\"otp\" name=\"otp\"><br><br>\n    <input type=\"hidden\" id=\"trans\" name=\"trans\" value={{.trans}}><br><br>\n    <input type=\"submit\" value=\"Submit\">\n</form>\n<p>Click the \"Submit\" button and the form-data will be sent to a page on th server called \"/form\".</p>\n</body>\n</html>"
			}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: invalid amount": {
			initAuthenticationMockCalled: false,
			givenBody: `{
							"number": "3301223454322203",
							"expired_date": "2023-03-22T00:00:00Z",
							"CVV": "999",
							"amount" : -100
						} `,
			expRs:       "{\"code\":\"invalid_request\", \"description\":\"invalid amount\"}",
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid number": {
			initAuthenticationMockCalled: false,
			givenBody: `{
							"number": "33012234",
							"expired_date": "2023-03-22T00:00:00Z",
							"CVV": "999",
							"amount" : 100
						} `,
			expRs:       "{\"code\":\"invalid_request\", \"description\":\"invalid number\"}",
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid CVV": {
			initAuthenticationMockCalled: false,
			givenBody: `{
							"number": "3301223454322203",
							"expired_date": "2023-03-22T00:00:00Z",
							"CVV": "999999",
							"amount" : 100
						} `,
			expRs:       "{\"code\":\"invalid_request\", \"description\":\"invalid CVV\"}",
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: invalid expired date": {
			initAuthenticationMockCalled: false,
			givenBody: `{
							"number": "3301223454322203",
							"expired_date": "2019-03-22T00:00:00Z",
							"CVV": "999",
							"amount" : 100
						} `,
			expRs:       "{\"code\":\"invalid_request\", \"description\":\"invalid expired date\"}",
			expHTTPCode: http.StatusBadRequest,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.initAuthenticationMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("InitAuthentication", mock.Anything, tc.initAuthentication.mockInCard, tc.initAuthentication.mockInOrder).
						Return(tc.initAuthentication.mockOutCard, tc.initAuthentication.mockOutOrder, tc.initAuthentication.mockErr),
				}
			}
			req := httptest.NewRequest(http.MethodPost, "/initAuthentication", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.InitAuthentication(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
