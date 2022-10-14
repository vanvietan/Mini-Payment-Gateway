package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// FindTransactionByID find transaction by ID
func (i impl) FindTransactionByID(ctx context.Context, transID int64) (model.Transaction, error) {
	trans, err := i.txRepo.FindTransactionByID(ctx, transID)
	if err != nil {
		log.Printf("error when find a transaction, %v", err)
		return model.Transaction{}, err
	}
	return trans, nil
}
