package transaction

import (
	"context"
	log "github.com/sirupsen/logrus"
	model2 "pg/api/internal/model"
)

// InitAuthentication check card in db and create an order
func (i impl) InitAuthentication(ctx context.Context, inputCard model2.Card, inputOrder model2.Order) (model2.Card, model2.Order, error) {

	/*
		find card existence
		create an order with amount
	*/

	card, err := i.cardRepo.GetCardByNumber(ctx, inputCard.Number)
	if err != nil {
		log.Printf("error when get card by number %v ", err)
		return model2.Card{}, model2.Order{}, err
	}
	//check amount and balance
	if card.Balance < 0 || card.Balance-inputOrder.Amount < 0 {
		return model2.Card{}, model2.Order{}, err
	}

	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return model2.Card{}, model2.Order{}, errG
	}
	inputOrder.ID = ID

	order, errO := i.orderRepo.CreateOrder(ctx, inputOrder)
	if errO != nil {
		log.Printf("error when create an order %v ", errO)
		return model2.Card{}, model2.Order{}, errO
	}

	return card, order, nil
}
