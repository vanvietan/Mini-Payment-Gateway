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

func TestFindTransactionByID(t *testing.T) {
	type findTransactionByID struct {
		mockID   int64
		mockResp model.Transaction
		mockErr  error
	}
	type arg struct {
		findTransactionByID findTransactionByID
		givenID             int64
		expRs               model.Transaction
		expErr              error
	}
	tcs := map[string]arg{
		"success: ": {
			findTransactionByID: findTransactionByID{
				mockID: 100,
				mockResp: model.Transaction{
					ID:        100,
					CardID:    100,
					OrderID:   100,
					OTP:       "123456",
					Status:    "PENDING",
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			givenID: 100,
			expRs: model.Transaction{
				ID:        100,
				CardID:    100,
				OrderID:   100,
				OTP:       "123456",
				Status:    "PENDING",
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: error from repo": {
			findTransactionByID: findTransactionByID{
				mockID:   100,
				mockResp: model.Transaction{},
				mockErr:  errors.New("something wrong"),
			},
			givenID: 100,
			expRs:   model.Transaction{},
			expErr:  errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			txRepo := new(mocks.Repository)
			cardRepo := new(mocksC.Repository)
			orderRepo := new(mocksO.Repository)
			txRepo.On("FindTransactionByID", mock.Anything, tc.findTransactionByID.mockID).
				Return(tc.findTransactionByID.mockResp, tc.findTransactionByID.mockErr)

			//WHEN
			svc := New(txRepo, cardRepo, orderRepo)
			rs, err := svc.FindTransactionByID(context.Background(), tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, rs)
			}
		})
	}
}
