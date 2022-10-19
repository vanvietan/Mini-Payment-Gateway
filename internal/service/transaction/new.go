package transaction

import (
	"context"
	"pg/internal/model"
	cardRepo "pg/internal/repository/card"
	orderRepo "pg/internal/repository/order"
	txRepo "pg/internal/repository/transaction"
)

// Service contains all service transaction
type Service interface {
	//CreateTransaction generate OTP and send back to clients
	CreateTransaction(ctx context.Context, cardID int64, orderID int64) (model.Transaction, error)
	
	//DeleteTransaction delete
	DeleteTransaction(ctx context.Context, transID int64) error

	//FindTransactionByID find a transaction
	FindTransactionByID(ctx context.Context, transID int64) (model.Transaction, error)

	//InitAuthentication check card in db and create an order
	InitAuthentication(ctx context.Context, inputCard model.Card, inputOrder model.Order) (model.Card, model.Order, error)

	//InitPayment init a payment
	InitPayment(ctx context.Context, transID int64) (model.Card, error)

	//AuthenticateTransaction authenticate transaction
	AuthenticateTransaction(ctx context.Context, id int64, otp string) error
}
type impl struct {
	txRepo    txRepo.Repository
	cardRepo  cardRepo.Repository
	orderRepo orderRepo.Repository
}

// New DI
func New(transaction txRepo.Repository, card cardRepo.Repository, order orderRepo.Repository) Service {
	return impl{
		txRepo:    transaction,
		cardRepo:  card,
		orderRepo: order,
	}
}
