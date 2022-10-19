package transaction

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	mocks "pg/internal/mocks/service/transaction"
	"testing"
)

func TestAuthenticatePayer(t *testing.T) {
	type authenticatePayer struct {
		mockOTP     string
		mockTransID int64
		mockErr     error
	}
	type arg struct {
		authenticatePayer           authenticatePayer
		authenticatePayerMockCalled bool
		givenID                     string
		givenOTP                    string
		expRs                       string
		expHTTPCode                 int
	}
	tcs := map[string]arg{
		"success: ": {
			authenticatePayer: authenticatePayer{
				mockOTP:     "123456",
				mockTransID: 100,
			},
			authenticatePayerMockCalled: true,
			givenID:                     "100",
			givenOTP:                    "123456",
			expRs:                       `{"message":"Successful Authenticated"}`,
			expHTTPCode:                 http.StatusOK,
		},
		"fail: invalid ID": {
			authenticatePayerMockCalled: false,
			givenID:                     "abc",
			givenOTP:                    "123456",
			expRs:                       `{"code":"invalid_request", "description":"id must be a number"}`,
			expHTTPCode:                 http.StatusBadRequest,
		},
		"fail: invalid OTP": {
			authenticatePayerMockCalled: false,
			givenID:                     "100",
			givenOTP:                    "abc",
			expRs:                       `{"code":"invalid_request", "description":"invalid OTP"}`,
			expHTTPCode:                 http.StatusBadRequest,
		},
		"fail: error from service ": {
			authenticatePayer: authenticatePayer{
				mockOTP:     "123456",
				mockTransID: 100,
				mockErr:     errors.New("something wrong"),
			},
			authenticatePayerMockCalled: true,
			givenID:                     "100",
			givenOTP:                    "123456",
			expRs:                       "{\"code\":\"internal_server_error\", \"description\":\"Incorrect OTP authenticate\"}",
			expHTTPCode:                 http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Mock
			mockSvc := new(mocks.Service)
			if tc.authenticatePayerMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("AuthenticateTransaction", mock.Anything, tc.authenticatePayer.mockTransID, tc.authenticatePayer.mockOTP).
						Return(tc.authenticatePayer.mockErr),
				}
			}
			// Given
			req := httptest.NewRequest(http.MethodPost, "/transactions/", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)
			form := url.Values{}
			form.Add("otp", tc.givenOTP)
			req.PostForm = form
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			// When
			instance := New(mockSvc)
			instance.AuthenticatePayer(res, req)
			// Then
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
