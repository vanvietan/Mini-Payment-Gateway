package transaction

import (
	"context"
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"pg/api/data"
	"pg/api/internal/model"
	"pg/api/internal/util"
	"testing"
	"time"
)

func TestUpdateTransaction(t *testing.T) {
	type arg struct {
		givenResult model.Transaction
		expResult   model.Transaction
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenResult: model.Transaction{
				ID:        101,
				CardID:    100,
				OrderID:   100,
				OTP:       "123456",
				Status:    "ACCEPTED",
				CreatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
			},
			expResult: model.Transaction{
				ID:        101,
				CardID:    100,
				OrderID:   100,
				OTP:       "123456",
				Status:    "ACCEPTED",
				CreatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"faiL: ": {
			givenResult: model.Transaction{
				ID:        101,
				CardID:    -100,
				OrderID:   100,
				OTP:       "123456",
				Status:    "ACCEPTED",
				CreatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 03, 14, 14, 0, 0, 0, time.UTC),
			},
			expResult: model.Transaction{},
			expErr:    errors.New("ERROR: insert or update on table \"transactions\" violates foreign key constraint \"transactions_card_id_fkey\" (SQLSTATE 23503)"),
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
			rs, err := instance.UpdateTransaction(context.Background(), tc.givenResult)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				if !cmp.Equal(tc.expResult, rs,
					cmpopts.IgnoreFields(model.Transaction{}, "CreatedAt", "UpdatedAt")) {
					t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, rs,
						cmp.Diff(tc.expResult, rs, cmpopts.IgnoreFields(model.Transaction{}, "CreatedAt", "UpdatedAt")))
					t.FailNow()
				}
			}
		})
	}
}
