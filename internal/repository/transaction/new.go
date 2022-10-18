package transaction

import (
	"context"
	"gorm.io/gorm"
	"pg/internal/model"
)

// Repository contains all transaction repository functions
type Repository interface {
	//CreateTransaction generate otp and create a transaction
	CreateTransaction(ctx context.Context, transaction model.Transaction) (model.Transaction, error)

	//FindTransactionByOTP compare the otp from clients with db
	FindTransactionByOTP(ctx context.Context, otp string) (model.Transaction, error)

	//FindTransactionByID find a transaction
	FindTransactionByID(ctx context.Context, transID int64) (model.Transaction, error)

	//UpdateTransaction update
	UpdateTransaction(ctx context.Context, input model.Transaction) (model.Transaction, error)

	//DeleteTransaction delete
	DeleteTransaction(ctx context.Context, transID int64) error
}
type impl struct {
	gormDB *gorm.DB
}

// New DI
func New(gormDB *gorm.DB) Repository {
	return impl{
		gormDB: gormDB,
	}
}
