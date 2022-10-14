package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// GetCardByID get card by its id
func (i impl) GetCardByID(ctx context.Context, cardID int64) (model.Card, error) {
	card, err := i.cardRepo.GetCardByID(ctx, cardID)
	if err != nil {
		log.Printf("error when get card by id , err: %v", err)
		return model.Card{}, err
	}
	return card, nil
}
