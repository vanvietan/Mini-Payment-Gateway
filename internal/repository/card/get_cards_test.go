package card

import (
	"context"
	"github.com/stretchr/testify/require"
	"pg/api/data"
	"pg/internal/model"
	"pg/internal/util"
	"testing"
	"time"
)

func TestGetCards(t *testing.T) {
	type arg struct {
		givenLimit int
		givenID    int64
		expResult  []model.Card
		expErr     error
	}
	tcs := map[string]arg{
		"success: get all cards": {
			givenLimit: 20,
			givenID:    0,
			expResult: []model.Card{
				{
					ID:          101,
					Number:      "3301223454322205",
					ExpiredDate: time.Date(2023, 3, 24, 00, 0, 0, 0, time.UTC),
					CVV:         "101",
					Balance:     10001,
					UserID:      101,
					CreatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
				},
				{
					ID:          100,
					Number:      "3301223454322204",
					ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
					CVV:         "100",
					Balance:     10000,
					UserID:      100,
					CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
				},
				{
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
		},
		"success: empty": {
			givenLimit: 3,
			givenID:    1,
			expResult:  []model.Card{},
		},
		"success: lastID is 101 ": {
			givenLimit: 3,
			givenID:    101,
			expResult: []model.Card{
				{
					ID:          100,
					Number:      "3301223454322204",
					ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
					CVV:         "100",
					Balance:     10000,
					UserID:      100,
					CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
				},
				{
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
		},
		"success: givenLimit 2": {
			givenLimit: 2,
			givenID:    0,
			expResult: []model.Card{
				{
					ID:          101,
					Number:      "3301223454322205",
					ExpiredDate: time.Date(2023, 3, 24, 00, 0, 0, 0, time.UTC),
					CVV:         "101",
					Balance:     10001,
					UserID:      101,
					CreatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 16, 14, 0, 0, 0, time.UTC),
				},
				{
					ID:          100,
					Number:      "3301223454322204",
					ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
					CVV:         "100",
					Balance:     10000,
					UserID:      100,
					CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
				},
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
			rs, err := instance.GetCards(context.Background(), tc.givenLimit, tc.givenID)

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
