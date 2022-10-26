package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"pg/api/internal/mocks/repository/order"
	"pg/api/internal/model"
	"pg/api/internal/util"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	type createOrder struct {
		mockIn   model.Order
		mockResp model.Order
		mockErr  error
	}
	type arg struct {
		givenIn     model.Order
		createOrder createOrder
		expRs       model.Order
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenIn: model.Order{
				ID:        103,
				Amount:    1000,
				CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			createOrder: createOrder{
				mockIn: model.Order{
					ID:        103,
					Amount:    1000,
					CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
				mockResp: model.Order{
					ID:        103,
					Amount:    1000,
					CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
			},
			expRs: model.Order{
				ID:        103,
				Amount:    1000,
				CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
		},
		"fail: error from repo": {
			givenIn: model.Order{
				ID:        103,
				Amount:    1000,
				CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			createOrder: createOrder{
				mockIn: model.Order{
					ID:        103,
					Amount:    1000,
					CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
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
			instance.On("CreateOrder", mock.Anything, tc.createOrder.mockIn).Return(tc.createOrder.mockResp, tc.createOrder.mockErr)

			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate id fail" {
					return 0, errors.New("something wrong")
				}
				return 103, nil
			}
			defer func() {
				getNextIDFunc = util.GetNextId
			}()

			//WHEN
			svc := New(instance)
			rs, err := svc.CreateOrder(context.Background(), tc.givenIn)

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
