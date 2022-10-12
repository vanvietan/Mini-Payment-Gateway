package transaction

import (
	"context"
	"pg/internal/repository/transaction"
)

// Service contains all service transaction
type Service interface {
	//GenerateOTP generate OTP and send back to clients
	GenerateOTP(ctx context.Context, cardID int64, orderID int64) (string, error)

	//CompareOTP compare OTP clients with db
	CompareOTP(ctx context.Context, input string) error

	//DeleteTransaction delete
	DeleteTransaction(ctx context.Context, transID int64) error
}
type impl struct {
	txRepo transaction.Repository
}

// New DI
func New(transaction transaction.Repository) Service {
	return impl{
		txRepo: transaction,
	}
}
