package card

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"pg/api/data"
	"pg/api/internal/model"
	"pg/api/internal/util"
	"testing"
	"time"
)

func TestUpdateCard(t *testing.T) {
	type arg struct {
		givenInput model.Card
		expResult  model.Card
		expErr     error
	}
	tcs := map[string]arg{
		"success: ": {
			givenInput: model.Card{
				ID:          101,
				Number:      "3301223454322205",
				ExpiredDate: time.Date(2023, 3, 24, 00, 0, 0, 0, time.UTC),
				CVV:         "101",
				Balance:     10001,
				UserID:      101,
				CreatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
			},
			expResult: model.Card{
				ID:          101,
				Number:      "3301223454322205",
				ExpiredDate: time.Date(2023, 3, 24, 00, 0, 0, 0, time.UTC),
				CVV:         "101",
				Balance:     10001,
				UserID:      101,
				CreatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
			},
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
			rs, err := instance.UpdateCard(context.Background(), tc.givenInput)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				if !cmp.Equal(tc.expResult, rs,
					cmpopts.IgnoreFields(model.Card{}, "CreatedAt", "UpdatedAt")) {
					t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, rs,
						cmp.Diff(tc.expResult, rs, cmpopts.IgnoreFields(model.Card{}, "CreatedAt", "UpdatedAt")))
					t.FailNow()
				}
			}
		})
	}
}
