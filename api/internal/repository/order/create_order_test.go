package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"pg/api/data"
	"pg/api/internal/model"
	"pg/api/internal/util"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	type arg struct {
		givenResult model.Order
		expResult   model.Order
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenResult: model.Order{
				ID:        103,
				Amount:    1000,
				CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expResult: model.Order{
				ID:        103,
				Amount:    1000,
				CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
		},
		"fail: error create: ": {
			givenResult: model.Order{
				ID:        103,
				Amount:    1000,
				CreatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expResult: model.Order{},
			expErr:    errors.New("ERROR: duplicate key value violates unique constraint \"orders_pkey\" (SQLSTATE 23505)"),
		},
	}
	dbConn, errDB := data.GetDatabaseConnection()
	require.NoError(t, errDB)

	errExe := util.ExecuteTestData(dbConn, "./testdata/order.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := New(dbConn)

			//WHEN
			rs, err := instance.CreateOrder(context.Background(), tc.givenResult)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rs)
			}

		})
	}
}
