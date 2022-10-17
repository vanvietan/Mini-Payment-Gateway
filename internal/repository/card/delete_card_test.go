package card

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"pg/api/data"
	"pg/internal/util"
	"testing"
)

func TestDeleteCard(t *testing.T) {
	type arg struct {
		givenID int64
		expErr  error
	}
	tcs := map[string]arg{
		"success: delete success": {
			givenID: 101,
		},
		"fail: no card id ": {
			givenID: 200,
			expErr:  errors.New("record not found"),
		},
	}

	dbConn, errDB := data.GetDatabaseConnection()
	require.NoError(t, errDB)

	errExe := util.ExecuteTestData(dbConn, "./testdata/card.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := New(dbConn)

			//WHEN
			err := instance.DeleteCard(context.Background(), tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}

		})
	}
}
