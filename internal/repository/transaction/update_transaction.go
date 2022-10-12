package transaction

import (
	"context"
	"pg/internal/model"
)

func (i impl) UpdateTransaction(ctx context.Context, input model.Transaction) (model.Transaction, error) {
	tx := i.gormDB.Save(&input)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return input, nil
}
