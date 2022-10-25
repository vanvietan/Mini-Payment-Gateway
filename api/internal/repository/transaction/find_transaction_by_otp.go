package transaction

import (
	"context"
	"pg/api/internal/model"
)

// FindTransactionByOTP compare OTP with clients( find by otp)
func (i impl) FindTransactionByOTP(ctx context.Context, otp string) (model.Transaction, error) {
	trans := model.Transaction{}
	tx := i.gormDB.Where("otp = ?", otp).First(&trans)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return trans, nil
}
