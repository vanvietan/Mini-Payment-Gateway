package transaction

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
)

// InitPayment init a payment
func (i impl) InitPayment(ctx context.Context, transID int64) (model.Card, error) {
	trans, err := i.txRepo.FindTransactionByID(ctx, transID)
	if err != nil {
		log.Printf("error when get transaction by id %v ", err)
		return model.Card{}, err
	}

	order, errO := i.orderRepo.GetOrderByID(ctx, trans.OrderID)
	if err != nil {
		log.Printf("error when get order by id %v ", errO)
		return model.Card{}, errO
	}

	card, errC := i.cardRepo.GetCardByID(ctx, trans.CardID)
	if errC != nil {
		log.Printf("error when find ID %v ", err)
		return model.Card{}, err
	}
	if card.Balance-order.Amount < 0 {
		return model.Card{}, errors.New("card balance is too low for this transaction")
	}
	card.Balance -= order.Amount

	cardUpdate, errU := i.cardRepo.UpdateCard(ctx, card)
	if errU != nil {
		log.Printf("error when save card %v ", err)
		return model.Card{}, err
	}

	errD := i.txRepo.DeleteTransaction(ctx, transID)
	if errD != nil {
		log.Printf("error when delete transaction %v ", errO)
		return model.Card{}, errD
	}

	return cardUpdate, nil
}
