package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"pg/api/internal/mocks/repository/order"
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	type deleteOrder struct {
		mockID  int64
		mockErr error
	}
	type arg struct {
		deleteOrder deleteOrder
		givenID     int64
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			deleteOrder: deleteOrder{
				mockID: 100,
			},
			givenID: 100,
		},
		"fail: error from repo": {
			deleteOrder: deleteOrder{
				mockID:  100,
				mockErr: errors.New("something wrong"),
			},
			givenID: 100,
			expErr:  errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("DeleteOrder", mock.Anything, tc.deleteOrder.mockID).
				Return(tc.deleteOrder.mockErr)

			//WHEN
			svc := New(instance)
			err := svc.DeleteOrder(context.Background(), tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
