package card

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// DeleteCard delete a card
func (i impl) DeleteCard(ctx context.Context, cardID int64) error {
	if err := i.cardRepo.DeleteCard(ctx, cardID); err != nil {
		log.Printf("error when delete a card %v", err)
		return err
	}
	return nil
}
