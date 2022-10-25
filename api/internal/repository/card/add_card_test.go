package card

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"pg/api/data"
	"pg/api/internal/model"
	"pg/api/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddCard(t *testing.T) {
	type arg struct {
		givenResult model.Card
		expResult   model.Card
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenResult: model.Card{
				ID:          103,
				Number:      "103",
				ExpiredDate: time.Date(2023, 4, 15, 16, 0, 0, 0, time.UTC),
				CVV:         "123456",
				Balance:     1000,
				UserID:      100,
				DeletedAt:   gorm.DeletedAt{},
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expResult: model.Card{
				ID:          103,
				Number:      "103",
				ExpiredDate: time.Date(2023, 4, 15, 16, 0, 0, 0, time.UTC),
				CVV:         "123456",
				Balance:     1000,
				UserID:      100,
				DeletedAt:   gorm.DeletedAt{},
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
		},
		"fail: error create": {
			givenResult: model.Card{
				ID:          103,
				Number:      "103",
				ExpiredDate: time.Date(2023, 4, 15, 16, 0, 0, 0, time.UTC),
				CVV:         "123456",
				Balance:     1000,
				UserID:      -100,
				DeletedAt:   gorm.DeletedAt{},
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expResult: model.Card{},
			expErr:    errors.New("ERROR: insert or update on table \"cards\" violates foreign key constraint \"cards_user_id_fkey\" (SQLSTATE 23503)"),
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
			rs, err := instance.AddCard(context.Background(), tc.givenResult)

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
