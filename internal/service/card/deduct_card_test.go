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

func TestDeductCard(t *testing.T) {
	type deductCard struct {
		mockID     int64
		mockAmount int64
		mockResp   model.Card
		mockErr    error
	}
	type arg struct {
		deductCard  deductCard
		givenID     int64
		givenAmount int64
		givenU      model.Card
		expRs       model.Card
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			deductCard: deductCard{
				mockID:     100,
				mockAmount: 100,
				mockResp: model.Card{
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
			givenID:     100,
			givenAmount: 100,
			expRs: model.Card{
				ID:          100,
				Number:      "3301223454322204",
				ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
				CVV:         "100",
				Balance:     9900,
				UserID:      100,
				CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
			},
			givenU: model.Card{
				ID:          100,
				Number:      "3301223454322204",
				ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
				CVV:         "100",
				Balance:     9900,
				UserID:      100,
				CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: error from repo ": {
			deductCard: deductCard{
				mockID:     100,
				mockAmount: 100,
				mockResp:   model.Card{},
				mockErr:    errors.New("something wrong"),
			},
			givenID:     100,
			givenAmount: 100,
			givenU: model.Card{
				ID:          100,
				Number:      "3301223454322204",
				ExpiredDate: time.Date(2023, 3, 23, 00, 0, 0, 0, time.UTC),
				CVV:         "100",
				Balance:     9900,
				UserID:      100,
				CreatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 15, 14, 0, 0, 0, time.UTC),
			},
			expRs:  model.Card{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("GetCardByID", mock.Anything, tc.deductCard.mockID).
				Return(tc.deductCard.mockResp, tc.deductCard.mockErr)
			instance.On("UpdateCard", mock.Anything, tc.givenU).
				Return(tc.expRs, tc.deductCard.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.DeductCard(context.Background(), tc.givenID, tc.givenAmount)

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
