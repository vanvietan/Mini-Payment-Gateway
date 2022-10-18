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
	card, err := i.cardSvc.GetCardByNumber(ctx, inputCard.Number)
	if err != nil {
		log.Printf("error when get card by number %v ", err)
		return model.Card{}, model.Order{}, err
	}
	order, errO := i.orderSvc.CreateOrder(ctx, inputOrder)
	if err != nil {
		log.Printf("error when create an order %v ", errO)
		return model.Card{}, model.Order{}, errO
	}

	return card, order, nil
}
