package transaction

import (
	"pg/internal/service/card"
	"pg/internal/service/order"
	"pg/internal/service/transaction"
)

type Handler struct {
	TxSvc    transaction.Service
	CardSvc  card.Service
	OrderSvc order.Service
}

// New DI
func New(txService transaction.Service, cardService card.Service, orderService order.Service) Handler {
	return Handler{
		TxSvc:    txService,
		CardSvc:  cardService,
		OrderSvc: orderService,
	}
}
