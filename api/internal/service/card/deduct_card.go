package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/api/internal/model"
)

// DeductCard deduct its balance
func (i impl) DeductCard(ctx context.Context, id int64, amount int64) (model.Card, error) {
	card, err := i.cardRepo.GetCardByID(ctx, id)
	if err != nil {
		log.Printf("error when find ID %v ", err)
		return model.Card{}, err
	}
	if card.Balance-amount < 0 {
		log.Printf("card balance is too low for this transaction")
		return model.Card{}, err
	}
	card.Balance -= amount
	cardUpdate, errU := i.cardRepo.UpdateCard(ctx, card)
	if errU != nil {
		log.Printf("error when save card %v ", err)
		return model.Card{}, err
	}
	return cardUpdate, nil
}
