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
	"testing"
)

func TestDeleteCard(t *testing.T) {
	type deleteCard struct {
		mockIn  int64
		mockErr error
	}
	type arg struct {
		deleteCard           deleteCard
		deleteCardMockCalled bool
		givenID              string
		expRs                string
		expHTTPCode          int
	}
	tcs := map[string]arg{
		"success": {
			deleteCard: deleteCard{
				mockIn: 99,
			},
			givenID:              "99",
			deleteCardMockCalled: true,
			expRs:                `{"message":"Deleted Card"}`,
			expHTTPCode:          http.StatusOK,
		},
		"fail: error from service": {
			deleteCard: deleteCard{
				mockIn:  99,
				mockErr: errors.New("something wrong"),
			},
			givenID:              "99",
			deleteCardMockCalled: true,
			expRs:                `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode:          http.StatusInternalServerError,
		},
		"fail: invalid ID": {
			deleteCardMockCalled: false,
			givenID:              "abc",
			expRs:                `{"code":"invalid_request", "description":"id must be a number"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.deleteCardMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("DeleteCard", mock.Anything, tc.deleteCard.mockIn).
						Return(tc.deleteCard.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodDelete, "/cards", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.DeleteCard(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
