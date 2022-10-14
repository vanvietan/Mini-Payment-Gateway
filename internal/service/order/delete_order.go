package order

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (i impl) DeleteOrder(ctx context.Context, id int64) error {
	if err := i.orderRepo.DeleteOrder(ctx, id); err != nil {
		log.Printf("error when delete a card %v", err)
		return err
	}
	return nil
}
