package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocksC "pg/internal/mocks/repository/card"
	mocksO "pg/internal/mocks/repository/order"
	mocks "pg/internal/mocks/repository/transaction"
	"pg/internal/model"
	"pg/internal/util"
	"testing"
	"time"
)

func TestInitAuthentication(t *testing.T) {
	type initAuthentication struct {
		mockCard      model.Card
		mockOrder     model.Order
		mockRespCard  model.Card
		mockRespOrder model.Order
		mockErr       error
	}
	type arg struct {
		initAuthentication initAuthentication
		givenCard          model.Card
		givenOrder         model.Order
		expRsCard          model.Card
		expRsOrder         model.Order
		expErr             error
	}

	tcs := map[string]arg{
		"success: ": {
			initAuthentication: initAuthentication{
				mockCard: model.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockRespCard: model.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockRespOrder: model.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			givenCard: model.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			givenOrder: model.Order{
				ID:        99,
				Amount:    99,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRsCard: model.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRsOrder: model.Order{
				ID:        99,
				Amount:    99,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: error from repo": {
			initAuthentication: initAuthentication{
				mockCard: model.Card{
					ID:          99,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model.Order{
					ID:        99,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockRespCard:  model.Card{},
				mockRespOrder: model.Order{},
				mockErr:       errors.New("something wrong"),
			},
			givenCard: model.Card{
				ID:          99,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9999,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			givenOrder: model.Order{
				ID:        99,
				Amount:    99,
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
			expRsCard:  model.Card{},
			expRsOrder: model.Order{},
			expErr:     errors.New("something wrong"),
		},
		"fail: generate id fail": {
			initAuthentication: initAuthentication{
				mockCard:      model.Card{},
				mockOrder:     model.Order{},
				mockRespCard:  model.Card{},
				mockRespOrder: model.Order{},
				mockErr:       errors.New("something wrong"),
			},
			givenCard:  model.Card{},
			givenOrder: model.Order{},
			expRsCard:  model.Card{},
			expRsOrder: model.Order{},
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
