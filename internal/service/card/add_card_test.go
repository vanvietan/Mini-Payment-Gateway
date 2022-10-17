package card

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "pg/internal/mocks/repository/card"
	"pg/internal/model"
	"pg/internal/util"
	"testing"
	"time"
)

func TestAddCard(t *testing.T) {
	type addCard struct {
		mockIn   model.Card
		mockResp model.Card
		mockErr  error
	}
	type arg struct {
		addCard    addCard
		givenInput model.Card
		expRs      model.Card
		expErr     error
	}
	tcs := map[string]arg{
		"success: ": {
			addCard: addCard{
				mockIn: model.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
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
			givenInput: model.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
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
		"fail: generate id fail": {
			addCard: addCard{
				mockIn: model.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockResp: model.Card{},
				mockErr:  errors.New("something wrong"),
			},
			givenInput: model.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRs:  model.Card{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("GetCardByNumber", mock.Anything, tc.addCard.mockIn.Number).Return(tc.addCard.mockResp, tc.addCard.mockErr)
			instance.On("AddCard", mock.Anything, tc.addCard.mockIn).Return(tc.addCard.mockResp, tc.addCard.mockErr)

			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate id fail" {
					return 0, errors.New("something wrong")
				}
				return 100, nil
			}
			defer func() {
				getNextIDFunc = util.GetNextId
			}()

			//WHEN
			svc := New(instance)
			rs, err := svc.AddCard(context.Background(), tc.givenInput)

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
