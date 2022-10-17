package card

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "pg/internal/mocks/repository/card"
	"pg/internal/model"
	"testing"
	"time"
)

func TestGetCards(t *testing.T) {
	type getCards struct {
		mockLimit int
		mockIn    int64
		mockResp  []model.Card
		mockErr   error
	}
	type arg struct {
		givenLimit  int
		givenLastID int64
		getCards    getCards
		expRs       []model.Card
		expErr      error
	}
	tcs := map[string]arg{
		"success: get cards": {
			givenLimit:  3,
			givenLastID: 0,
			getCards: getCards{
				mockLimit: 3,
				mockIn:    0,
				mockResp: []model.Card{
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
			expRs: []model.Card{
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
		"fail: empty result": {
			givenLimit:  3,
			givenLastID: 1,
			getCards: getCards{
				mockLimit: 3,
				mockIn:    1,
				mockResp:  []model.Card{},
				mockErr:   errors.New("something wrong"),
			},
			expRs:  []model.Card{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("GetCards", mock.Anything, tc.givenLimit, tc.givenLastID).Return(tc.getCards.mockResp, tc.getCards.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetCards(context.Background(), tc.givenLimit, tc.givenLastID)

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
