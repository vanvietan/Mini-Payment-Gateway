package card

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"pg/api/internal/mocks/repository/card"
	"testing"
)

func TestDeleteCard(t *testing.T) {
	type deleteCard struct {
		mockIn  int64
		mockErr error
	}
	type arg struct {
		givenIn    int64
		deleteCard deleteCard
		expErr     error
	}
	tcs := map[string]arg{
		"success: ": {
			givenIn: 99,
			deleteCard: deleteCard{
				mockIn: 99,
			},
		},
		"fail: error from repo": {
			givenIn: 200,
			deleteCard: deleteCard{
				mockIn:  200,
				mockErr: errors.New("something wrong"),
			},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.Repository)
			instance.On("DeleteCard", mock.Anything, tc.deleteCard.mockIn).
				Return(tc.deleteCard.mockErr)

			//WHEN
			svc := New(instance)
			err := svc.DeleteCard(context.Background(), tc.givenIn)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
