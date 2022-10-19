package handler

import (
	"pg/internal/handler/card"
	"pg/internal/handler/transaction"
	cardService "pg/internal/service/card"
	orderService "pg/internal/service/order"
	transactionService "pg/internal/service/transaction"
)

// Handler DI
type Handler struct {
	CardHandler  card.Handler
	TransHandler transaction.Handler
}

// New DI
func New(cardSvc cardService.Service, transactionSvc transactionService.Service, orderSvc orderService.Service) Handler {
	return Handler{
		CardHandler: card.Handler{
			CardSvc: cardSvc,
		},
		TransHandler: transaction.Handler{
			TxSvc: transactionSvc,
		},
	}
}
