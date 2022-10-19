package transaction

import (
	"context"
	"pg/internal/model"
)

// CreateTransaction generate OTP and create a authentication
func (i impl) CreateTransaction(ctx context.Context, transaction model.Transaction) (model.Transaction, error) {

	tx := i.gormDB.Create(&transaction)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return transaction, nil
}
