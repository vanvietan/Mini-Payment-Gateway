package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

func (i impl) UpdateCard(ctx context.Context, input model.Card, cardID int64) (model.Card, error) {
	cardF, err := i.cardRepo.GetCardByID(ctx, cardID)
	if err != nil {
		log.Printf("error when find an card by ID: %d", cardID)
		return model.Card{}, err
	}
	cardF.ID = input.ID
	cardF.ExpiredDate = input.ExpiredDate
	cardF.CVV = input.CVV
	cardF.Balance = input.Balance

	cardU, err := i.cardRepo.UpdateCard(ctx, cardF)
	if err != nil {
		log.Printf("error when save card %+v", input)
		return model.Card{}, err
	}
	return cardU, nil
}
