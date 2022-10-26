package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/api/internal/model"
)

// GetCards get all cards
func (i impl) GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error) {
	cards, err := i.cardRepo.GetCards(ctx, limit, lastID)
	if err != nil {
		log.Printf("error when get orders, limit %d, last %d", limit, lastID)
		return nil, err
	}
	return cards, nil
}
