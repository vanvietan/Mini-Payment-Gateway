package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// InitAuthentication check card in db and create an order
func (i impl) InitAuthentication(ctx context.Context, inputCard model.Card, inputOrder model.Order) (model.Card, model.Order, error) {

	/*
		find card existence
		create an order with amount
	*/

	card, err := i.cardRepo.GetCardByNumber(ctx, inputCard.Number)
	if err != nil {
		log.Printf("error when get card by number %v ", err)
		return model.Card{}, model.Order{}, err
	}

	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return model.Card{}, model.Order{}, errG
	}
	inputOrder.ID = ID

	order, errO := i.orderRepo.CreateOrder(ctx, inputOrder)
	if errO != nil {
		log.Printf("error when create an order %v ", errO)
		return model.Card{}, model.Order{}, errO
	}

	return card, order, nil
}
