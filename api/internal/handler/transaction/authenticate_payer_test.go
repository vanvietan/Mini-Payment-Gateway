package transaction

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	mocks "pg/api/internal/mocks/service/transaction"
	"pg/api/internal/model"
	"testing"
)

func TestAuthenticatePayer(t *testing.T) {
	type authenticatePayer struct {
		mockInCardID  int64
		mockInOrderID int64
		mockResp      model.Transaction
		mockErr       error
	}
	type arg struct {
		authenticatePayer           authenticatePayer
		authenticatePayerMockCalled bool
		givenCardID                 string
		givenOrderID                string
		expRs                       string
		expHTTPCode                 int
	}
	tcs := map[string]arg{
		"success: ": {
			authenticatePayer: authenticatePayer{
				mockInCardID:  100,
				mockInOrderID: 100,
				mockResp: model.Transaction{
					ID:      100,
					CardID:  100,
					OrderID: 100,
					OTP:     "123456",
					Status:  "PENDING",
				},
			},
			authenticatePayerMockCalled: true,
			givenCardID:                 "100",
			givenOrderID:                "100",
			expRs:                       `<!DOCTYPE html>\n<html>\n<body>\n<h1>Submit your OTP</h1>\n<form action=\"/authenticateTransaction/100\" method=\"post\">\n    <label for=\"otp\">OTP:</label>\n    <input type=\"text\" id=\"otp\" name=\"otp\"><br><br>\n    <input type=\"submit\" value=\"Submit\">\n</form>\n</body>\n</html>`,
			expHTTPCode:                 http.StatusOK,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Mock
			mockSvc := new(mocks.Service)
			if tc.authenticatePayerMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("CreateTransaction", mock.Anything, tc.authenticatePayer.mockInCardID, tc.authenticatePayer.mockInOrderID).
						Return(tc.authenticatePayer.mockResp, tc.authenticatePayer.mockErr),
				}
			}
			// Given
			req := httptest.NewRequest(http.MethodPost, "/authenticatePayer/", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("card", tc.givenCardID)
			routeCtx.URLParams.Add("order", tc.givenOrderID)

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			// When
			instance := New(mockSvc)
			instance.AuthenticateTransaction(res, req)
			// Then
			//require.Equal(t, tc.expHTTPCode, res.Code)
			//require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
