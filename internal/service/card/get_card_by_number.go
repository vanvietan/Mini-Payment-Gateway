package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

func (i impl) GetCardByNumber(ctx context.Context, cardNumber string) (model.Card, error) {
	card, err := i.cardRepo.GetCardByNumber(ctx, cardNumber)
	if err != nil {
		log.Printf("error when get card by number , err: %v", err)
		return model.Card{}, err
	}
	return card, nil
}
