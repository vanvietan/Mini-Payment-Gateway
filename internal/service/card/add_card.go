package card

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
	"pg/internal/util"
)

var getNextIDFunc = util.GetNextId

// AddCard add a card
func (i impl) AddCard(ctx context.Context, input model.Card) (model.Card, error) {

	cardN, _ := i.GetCardByNumber(ctx, input.Number)
	if (cardN.Number) == input.Number {
		return cardN, nil
	}

	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return model.Card{}, err
	}
	input.ID = ID

	card, errI := i.cardRepo.AddCard(ctx, input)
	if errI != nil {
		log.Printf("error when add a card: %+v", input)
		return model.Card{}, err
	}
	return card, nil
}
