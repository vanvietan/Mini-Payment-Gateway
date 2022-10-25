package card

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

func TestGetCardByNumber(t *testing.T) {
	type arg struct {
		givenString string
		expResult   model.Card
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenString: "3301223454322203",
			expResult: model.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: can't find a card": {
			givenString: "123",
			expResult:   model.Card{},
			expErr:      errors.New("record not found"),
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
			rs, err := instance.GetCardByNumber(context.Background(), tc.givenString)

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
