package card

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"pg/api/internal/mocks/repository/card"
	"pg/api/internal/model"
	"testing"
	"time"
)

func TestGetCardByID(t *testing.T) {
	type getCardByID struct {
		mockIn   int64
		mockResp model.Card
		mockErr  error
	}
	type arg struct {
		givenID     int64
		getCardByID getCardByID
		expRs       model.Card
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: 99,
			getCardByID: getCardByID{
				mockIn: 99,
				mockResp: model.Card{
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
			expRs: model.Card{
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
		"fail: can't find the id": {
			givenID: 200,
			getCardByID: getCardByID{
				mockIn:   200,
				mockResp: model.Card{},
				mockErr:  errors.New("something wrong"),
			},
			expRs:  model.Card{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("GetCardByID", mock.Anything, tc.getCardByID.mockIn).Return(tc.getCardByID.mockResp, tc.getCardByID.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetCardByID(context.Background(), tc.givenID)

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
