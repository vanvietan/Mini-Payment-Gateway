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
	"testing"
	"time"
)

func TestInitPayment(t *testing.T) {
	type initPayment struct {
		mockTransID int64
		mockTrans   model2.Transaction
		mockCard    model2.Card
		mockOrder   model2.Order
		mockResp    model2.Card
		mockErr     error
	}
	type arg struct {
		initPayment  initPayment
		givenTransID int64
		expRs        model2.Card
		expErr       error
	}

	tcs := map[string]arg{
		"success: ": {
			initPayment: initPayment{
				mockTransID: 100,
				mockTrans: model2.Transaction{
					ID:      100,
					CardID:  100,
					OrderID: 100,
					OTP:     "123456",
					Status:  model2.StatusAccepted,
				},
				mockCard: model2.Card{
					ID:          100,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model2.Order{
					ID:        100,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockResp: model2.Card{
					ID:          100,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9900,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			givenTransID: 100,
			expRs: model2.Card{
				ID:          100,
				Number:      "3301223454322203",
				ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
				CVV:         "999",
				Balance:     9900,
				UserID:      99,
				CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail:user has not authenticated ": {
			initPayment: initPayment{
				mockTransID: 100,
				mockTrans: model2.Transaction{
					ID:      100,
					CardID:  100,
					OrderID: 100,
					OTP:     "123456",
					Status:  model2.StatusPending,
				},
				mockCard: model2.Card{
					ID:          100,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model2.Order{
					ID:        100,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockResp: model2.Card{},
				mockErr:  errors.New("something wrong"),
			},
			givenTransID: 100,
			expRs:        model2.Card{},
			expErr:       errors.New("something wrong"),
		},
		"fail: error from repo": {
			initPayment: initPayment{
				mockTransID: 100,
				mockTrans: model2.Transaction{
					ID:      100,
					CardID:  100,
					OrderID: 100,
					OTP:     "123456",
					Status:  model2.StatusAccepted,
				},
				mockCard: model2.Card{
					ID:          100,
					Number:      "3301223454322203",
					ExpiredDate: time.Date(2023, 3, 22, 00, 0, 0, 0, time.UTC),
					CVV:         "999",
					Balance:     9999,
					UserID:      99,
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockOrder: model2.Order{
					ID:        100,
					Amount:    99,
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockResp: model2.Card{},
				mockErr:  errors.New("something wrong"),
			},
			givenTransID: 100,
			expRs:        model2.Card{},
			expErr:       errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			txRepo := new(mocks.Repository)
			cardRepo := new(mocksC.Repository)
			orderRepo := new(mocksO.Repository)
			txRepo.On("FindTransactionByID", mock.Anything, tc.initPayment.mockTransID).
				Return(tc.initPayment.mockTrans, tc.initPayment.mockErr)
			orderRepo.On("GetOrderByID", mock.Anything, tc.initPayment.mockTrans.OrderID).
				Return(tc.initPayment.mockOrder, tc.initPayment.mockErr)
			cardRepo.On("GetCardByID", mock.Anything, tc.initPayment.mockTrans.CardID).
				Return(tc.initPayment.mockCard, tc.initPayment.mockErr)
			cardRepo.On("UpdateCard", mock.Anything, tc.initPayment.mockResp).
				Return(tc.initPayment.mockResp, tc.initPayment.mockErr)
			txRepo.On("DeleteTransaction", mock.Anything, tc.initPayment.mockTransID).
				Return(tc.initPayment.mockErr)

			//WHEN
			svc := New(txRepo, cardRepo, orderRepo)
			rs, err := svc.InitPayment(context.Background(), tc.givenTransID)

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
