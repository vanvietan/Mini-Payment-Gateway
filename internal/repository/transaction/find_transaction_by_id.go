package transaction

import (
	"context"
	"pg/internal/model"
)

func (i impl) FindTransactionByID(ctx context.Context, transID int64) (model.Transaction, error) {
	trans := model.Transaction{}
	tx := i.gormDB.First(&trans, transID)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return trans, nil
}
