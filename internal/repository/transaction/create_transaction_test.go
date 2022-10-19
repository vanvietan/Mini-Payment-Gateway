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

func TestGenerateOTP(t *testing.T) {
	type arg struct {
		givenResult model.Transaction
		expResult   string
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenResult: model.Transaction{
				ID:        101,
				CardID:    100,
				OrderID:   100,
				Status:    "PENDING",
				CreatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
			},
			expResult: "512369",
		},
		"fail: ": {
			givenResult: model.Transaction{
				ID:        100,
				CardID:    100,
				OrderID:   100,
				Status:    "PENDING",
				CreatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
			},
			expResult: "",
			expErr:    errors.New("ERROR: duplicate key value violates unique constraint \"transactions_pkey\" (SQLSTATE 23505)"),
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
			rs, err := instance.CreateTransaction(context.Background(), tc.givenResult)

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
