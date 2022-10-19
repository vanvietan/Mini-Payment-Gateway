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
	"pg/internal/util"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	type createTransaction struct {
		mockCard  int64
		mockOrder int64
		mockTrans model.Transaction
		mockResp  model.Transaction
		mockErr   error
	}
	type arg struct {
		givenCardID       int64
		givenOrderID      int64
		createTransaction createTransaction
		expRs             model.Transaction
		expErr            error
	}
	tcs := map[string]arg{
		"success: ": {
			givenCardID:  100,
			givenOrderID: 100,
			createTransaction: createTransaction{
				mockCard:  100,
				mockOrder: 100,
				mockTrans: model.Transaction{
					ID:      100,
					CardID:  100,
					OrderID: 100,
					OTP:     "123456",
					Status:  "PENDING",
				},
				mockResp: model.Transaction{
					ID:      100,
					CardID:  100,
					OrderID: 100,
					OTP:     "123456",
					Status:  "PENDING",
				},
			},
			expRs: model.Transaction{
				ID:      100,
				CardID:  100,
				OrderID: 100,
				OTP:     "123456",
				Status:  "PENDING",
			},
		},
		"fail: generate id fail": {
			givenCardID:  100,
			givenOrderID: 100,
			createTransaction: createTransaction{
				mockCard:  100,
				mockOrder: 100,
				mockTrans: model.Transaction{},
				mockResp:  model.Transaction{},
				mockErr:   errors.New("something wrong"),
			},
			expRs:  model.Transaction{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			txRepo := new(mocks.Repository)
			cardRepo := new(mocksC.Repository)
			orderRepo := new(mocksO.Repository)
			txRepo.On("CreateTransaction", mock.Anything, tc.createTransaction.mockTrans).
				Return(tc.createTransaction.mockResp, tc.createTransaction.mockErr)

			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate id fail" {
					return 0, errors.New("something wrong")
				}
				return 100, nil
			}
			randomFunc = func(min int, max int) int {
				return 123456
			}
			defer func() {
				randomFunc = randInt
				getNextIDFunc = util.GetNextId
			}()

			//WHEN
			svc := New(txRepo, cardRepo, orderRepo)
			rs, err := svc.CreateTransaction(context.Background(), tc.givenCardID, tc.givenOrderID)

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
