package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocksC "pg/api/internal/mocks/repository/card"
	mocksO "pg/api/internal/mocks/repository/order"
	"pg/api/internal/mocks/repository/transaction"
	model2 "pg/api/internal/model"
	"pg/api/internal/util"
	"testing"
	"time"
)

func TestInitAuthentication(t *testing.T) {
	type initAuthentication struct {
		mockCard      model2.Card
		mockOrder     model2.Order
		mockRespCard  model2.Card
		mockRespOrder model2.Order
		mockErr       error
	}
	type arg struct {
		initAuthentication initAuthentication
		givenCard          model2.Card
		givenOrder         model2.Order
		expRsCard          model2.Card
		expRsOrder         model2.Order
		expErr             error
	}

	tcs := map[string]arg{
		"success: ": {
			initAuthentication: initAuthentication{
				mockCard: model2.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model2.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockRespCard: model2.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockRespOrder: model2.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			givenCard: model2.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			givenOrder: model2.Order{
				ID:        99,
				Amount:    99,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRsCard: model2.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRsOrder: model2.Order{
				ID:        99,
				Amount:    99,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: error from repo": {
			initAuthentication: initAuthentication{
				mockCard: model2.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model2.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockRespCard:  model2.Card{},
				mockRespOrder: model2.Order{},
				mockErr:       errors.New("something wrong"),
			},
			givenCard: model2.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			givenOrder: model2.Order{
				ID:        99,
				Amount:    99,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRsCard:  model2.Card{},
			expRsOrder: model2.Order{},
			expErr:     errors.New("something wrong"),
		},
		"fail: generate id fail": {
			initAuthentication: initAuthentication{
				mockCard:      model2.Card{},
				mockOrder:     model2.Order{},
				mockRespCard:  model2.Card{},
				mockRespOrder: model2.Order{},
				mockErr:       errors.New("something wrong"),
			},
			givenCard:  model2.Card{},
			givenOrder: model2.Order{},
			expRsCard:  model2.Card{},
			expRsOrder: model2.Order{},
			expErr:     errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			txRepo := new(mocks.Repository)
			cardRepo := new(mocksC.Repository)
			orderRepo := new(mocksO.Repository)
			cardRepo.On("GetCardByNumber", mock.Anything, tc.initAuthentication.mockCard.Number).
				Return(tc.initAuthentication.mockCard, tc.initAuthentication.mockErr)
			orderRepo.On("CreateOrder", mock.Anything, tc.initAuthentication.mockOrder).
				Return(tc.initAuthentication.mockOrder, tc.initAuthentication.mockErr)

			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate id fail" {
					return 0, errors.New("something wrong")
				}
				return 99, nil
			}
			//monkey patching
			defer func() {
				getNextIDFunc = util.GetNextId
			}()

			//WHEN
			svc := New(txRepo, cardRepo, orderRepo)
			rsC, rsO, err := svc.InitAuthentication(context.Background(), tc.givenCard, tc.givenOrder)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRsOrder, rsO)
				require.Equal(t, tc.expRsCard, rsC)
			}
		})
	}
}
