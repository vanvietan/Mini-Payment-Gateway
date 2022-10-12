package transaction

import (
	"context"
	"errors"
	"pg/internal/model"
)

// CompareOTP compare OTP with clients
func (i impl) CompareOTP(ctx context.Context, otp string) (model.Transaction, error) {
	trans := model.Transaction{}
	tx := i.gormDB.Where("otp = ?", otp).Find(&trans)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	if tx.RowsAffected != 1 {
		return model.Transaction{}, errors.New("record not found")
	}

	return trans, nil
}
