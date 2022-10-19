package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"pg/api/data"
	"pg/internal/model"
	"pg/internal/util"
	"testing"
	"time"
)

func TestFindTransactionByID(t *testing.T) {
	type arg struct {
		givenID   int64
		expResult model.Transaction
		expErr    error
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: 100,
			expResult: model.Transaction{
				ID:        100,
				CardID:    100,
				OrderID:   100,
				OTP:       "123456",
				Status:    "PENDING",
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail:record not found": {
			givenID:   200,
			expResult: model.Transaction{},
			expErr:    errors.New("record not found"),
		},
	}
	dbConn, errDB := data.GetDatabaseConnection()
	require.NoError(t, errDB)

	errExe := util.ExecuteTestData(dbConn, "./testdata/transaction.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := New(dbConn)

			//WHEN
			rs, err := instance.FindTransactionByID(context.Background(), tc.givenID)

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
