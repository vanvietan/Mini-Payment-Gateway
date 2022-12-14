package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocksC "pg/api/internal/mocks/repository/card"
	mocksO "pg/api/internal/mocks/repository/order"
	"pg/api/internal/mocks/repository/transaction"
	"testing"
)

func TestDeleteTransaction(t *testing.T) {
	type deleteTransaction struct {
		mockIn  int64
		mockErr error
	}
	type arg struct {
		givenIn           int64
		deleteTransaction deleteTransaction
		expErr            error
	}
	tcs := map[string]arg{
		"success: ": {
			givenIn: 99,
			deleteTransaction: deleteTransaction{
				mockIn: 99,
			},
		},
		"fail: error from repo": {
			givenIn: 200,
			deleteTransaction: deleteTransaction{
				mockIn:  200,
				mockErr: errors.New("something wrong"),
			},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			txRepo := new(mocks.Repository)
			cardRepo := new(mocksC.Repository)
			orderRepo := new(mocksO.Repository)
			txRepo.On("DeleteTransaction", mock.Anything, tc.deleteTransaction.mockIn).
				Return(tc.deleteTransaction.mockErr)

			//WHEN
			svc := New(txRepo, cardRepo, orderRepo)
			err := svc.DeleteTransaction(context.Background(), tc.givenIn)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
