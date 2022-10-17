package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "pg/internal/mocks/repository/transaction"
	"pg/internal/model"
	"testing"
	"time"
)

func TestCompareOTP(t *testing.T) {
	type compareOTP struct {
		mockIn    string
		mockTrans model.Transaction
		mockResp  model.Transaction
		mockErr   error
	}
	type arg struct {
		compareOTP compareOTP
		givenIn    string
		expRs      model.Transaction
		expErr     error
	}
	tcs := map[string]arg{
		"success: ": {
			compareOTP: compareOTP{
				mockIn: "123456",
				mockTrans: model.Transaction{
					ID:        100,
					CardID:    100,
					OrderID:   100,
					OTP:       "123456",
					Status:    "ACCEPTED",
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
				mockResp: model.Transaction{
					ID:        100,
					CardID:    100,
					OrderID:   100,
					OTP:       "123456",
					Status:    "ACCEPTED",
					CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
			givenIn: "123456",
			expRs: model.Transaction{
				ID:        100,
				CardID:    100,
				OrderID:   100,
				OTP:       "123456",
				Status:    "ACCEPTED",
				CreatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
			},
		},
		"fail: error when compare": {
			compareOTP: compareOTP{
				mockIn:    "123456",
				mockTrans: model.Transaction{},
				mockResp:  model.Transaction{},
				mockErr:   errors.New("something wrong"),
			},
			givenIn: "123456",
			expRs:   model.Transaction{},
			expErr:  errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("CompareOTP", mock.Anything, tc.compareOTP.mockIn).
				Return(tc.compareOTP.mockResp, tc.compareOTP.mockErr)
			instance.On("UpdateTransaction", mock.Anything, tc.compareOTP.mockTrans).
				Return(tc.compareOTP.mockResp, tc.compareOTP.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.CompareOTP(context.Background(), tc.givenIn)

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
