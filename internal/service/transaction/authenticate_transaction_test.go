package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocksC "pg/internal/mocks/repository/card"
	mocksO "pg/internal/mocks/repository/order"
	mocks "pg/internal/mocks/repository/transaction"
	"pg/internal/model"
	"testing"
	"time"
)

func TestAuthenticateTransaction(t *testing.T) {
	type authenticateTransaction struct {
		mockID    int64
		mockOTP   string
		mockTrans model.Transaction
		mockErr   error
	}
	type arg struct {
		authenticateTransaction authenticateTransaction
		givenID                 int64
		givenOTP                string
		expErr                  error
	}
	tcs := map[string]arg{
		"success: ": {
			authenticateTransaction: authenticateTransaction{
				mockID:  100,
				mockOTP: "123456",
				mockTrans: model.Transaction{
					ID:        100,
					CardID:    100,
					OrderID:   100,
					OTP:       "123456",
					Status:    "ACCEPTED",
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			givenID:  100,
			givenOTP: "123456",
		},
		"fail: error from repo": {
			authenticateTransaction: authenticateTransaction{
				mockID:  100,
				mockOTP: "123456",
				mockTrans: model.Transaction{
					ID:        100,
					CardID:    100,
					OrderID:   100,
					OTP:       "123456",
					Status:    "ACCEPTED",
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockErr: errors.New("something wrong"),
			},
			givenID:  100,
			givenOTP: "123456",
			expErr:   errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			txRepo := new(mocks.Repository)
			cardRepo := new(mocksC.Repository)
			orderRepo := new(mocksO.Repository)
			txRepo.On("FindTransactionByID", mock.Anything, tc.authenticateTransaction.mockID).
				Return(tc.authenticateTransaction.mockTrans, tc.authenticateTransaction.mockErr)
			txRepo.On("UpdateTransaction", mock.Anything, tc.authenticateTransaction.mockTrans).
				Return(tc.authenticateTransaction.mockTrans, tc.authenticateTransaction.mockErr)

			//WHEN
			svc := New(txRepo, cardRepo, orderRepo)
			err := svc.AuthenticateTransaction(context.Background(), tc.givenID, tc.givenOTP)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
