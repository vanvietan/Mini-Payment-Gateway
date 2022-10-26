package handler

import (
	"pg/api/internal/handler/card"
	"pg/api/internal/handler/transaction"
	cardService "pg/api/internal/service/card"
	orderService "pg/api/internal/service/order"
	transactionService "pg/api/internal/service/transaction"
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
