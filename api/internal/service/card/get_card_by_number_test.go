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

func TestGetCardByNumber(t *testing.T) {
	type getCardByNumber struct {
		mockIn   string
		mockResp model.Card
		mockErr  error
	}
	type arg struct {
		givenIn         string
		getCardByNumber getCardByNumber
		expRs           model.Card
		expErr          error
	}
	tcs := map[string]arg{
		"success: ": {
			givenIn: "3301223454322203",
			getCardByNumber: getCardByNumber{
				mockIn: "3301223454322203",
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
		"fail: can't find the number": {
			givenIn: "123",
			getCardByNumber: getCardByNumber{
				mockIn:   "123",
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
			instance.On("GetCardByNumber", mock.Anything, tc.getCardByNumber.mockIn).Return(tc.getCardByNumber.mockResp, tc.getCardByNumber.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetCardByNumber(context.Background(), tc.givenIn)

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
