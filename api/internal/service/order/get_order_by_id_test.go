package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"pg/api/internal/mocks/repository/order"
	"pg/api/internal/model"
	"testing"
	"time"
)

func TestGetOrderByID(t *testing.T) {
	type getOrderByID struct {
		mockIn   int64
		mockResp model.Order
		mockErr  error
	}
	type arg struct {
		givenID      int64
		getOrderByID getOrderByID
		expRs        model.Order
		expErr       error
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: 100,
			getOrderByID: getOrderByID{
				mockIn: 100,
				mockResp: model.Order{
					ID:        100,
					Amount:    1000,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			expRs: model.Order{
				ID:        100,
				Amount:    1000,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: error from repo": {
			givenID: 0,
			getOrderByID: getOrderByID{
				mockIn:   0,
				mockResp: model.Order{},
				mockErr:  errors.New("something wrong"),
			},
			expRs:  model.Order{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("GetOrderByID", mock.Anything, tc.getOrderByID.mockIn).
				Return(tc.getOrderByID.mockResp, tc.getOrderByID.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetOrderByID(context.Background(), tc.givenID)

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
