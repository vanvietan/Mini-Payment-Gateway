package transaction

import (
	"context"
	"gorm.io/gorm"
	"pg/internal/model"
)

// Repository contains all transaction repository functions
type Repository interface {
	//GenerateOTP generate otp and create a transaction
	GenerateOTP(ctx context.Context, transaction model.Transaction) (string, error)
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
