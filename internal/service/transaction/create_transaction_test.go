package transaction

//import (
//	"context"
//	"errors"
//	"github.com/stretchr/testify/mock"
//	"github.com/stretchr/testify/require"
//	mocks "pg/internal/mocks/repository/transaction"
//	"pg/internal/model"
//	"pg/internal/util"
//	"testing"
//)
//
//func TestGenerateOTP(t *testing.T) {
//	type generateOTP struct {
//		mockCard  int64
//		mockOrder int64
//		mockTrans model.Transaction
//		mockResp  string
//		mockErr   error
//	}
//	type arg struct {
//		givenCardID  int64
//		givenOrderID int64
//		generateOTP  generateOTP
//		expRs        string
//		expErr       error
//	}
//	tcs := map[string]arg{
//		"success: ": {
//			givenCardID:  100,
//			givenOrderID: 100,
//			generateOTP: generateOTP{
//				mockCard:  100,
//				mockOrder: 100,
//				mockTrans: model.Transaction{
//					ID:      100,
//					CardID:  100,
//					OrderID: 100,
//					OTP:     "",
//					Status:  "PENDING",
//				},
//				mockResp: "123456",
//			},
//			expRs: "123456",
//		},
//		"fail: generate id fail": {
//			givenCardID:  100,
//			givenOrderID: 100,
//			generateOTP: generateOTP{
//				mockCard:  100,
//				mockOrder: 100,
//				mockTrans: model.Transaction{},
//				mockResp:  "",
//				mockErr:   errors.New("something wrong"),
//			},
//			expRs:  "",
//			expErr: errors.New("something wrong"),
//		},
//	}
//	for s, tc := range tcs {
//		t.Run(s, func(t *testing.T) {
//			//GIVEN
//			instance := new(mocks.Repository)
//			instance.On("CreateTransaction", mock.Anything, tc.generateOTP.mockTrans).
//				Return(tc.generateOTP.mockResp, tc.generateOTP.mockErr)
//
//			getNextIDFunc = func() (int64, error) {
//				if s == "fail: generate id fail" {
//					return 0, errors.New("something wrong")
//				}
//				return 100, nil
//			}
//			defer func() {
//				getNextIDFunc = util.GetNextId
//			}()
//
//			//WHEN
//			svc := New(instance)
//			rs, err := svc.CreateTransaction(context.Background(), tc.givenCardID, tc.givenOrderID)
//
//			//THEN
//			if tc.expErr != nil {
//				require.EqualError(t, err, tc.expErr.Error())
//			} else {
//				require.NoError(t, err)
//				require.Equal(t, tc.expRs, rs)
//			}
//		})
//	}
//}
