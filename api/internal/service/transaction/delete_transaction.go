package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (i impl) DeleteTransaction(ctx context.Context, transID int64) error {
	if err := i.txRepo.DeleteTransaction(ctx, transID); err != nil {
		log.Printf("error when delete a transaction %v", err)
		return err
	}
	return nil
}
