package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// AddCard add a card
func (i impl) AddCard(ctx context.Context, input model.Card) (model.Card, error) {
	card, err := i.cardRepo.AddCard(ctx, input)
	if err != nil {
		log.Printf("error when add a card: %+v", input)
		return model.Card{}, err
	}
	return card, nil
}
