package order

import (
	"context"
	"gorm.io/gorm"
	"pg/internal/model"
)

// Repository contains all order repository functions
type Repository interface {
	//CreateOrder add an order
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)

	//GetOrderByID get an order by its id
	GetOrderByID(ctx context.Context, id int64) (model.Order, error)

	//DeleteOrder delete an order
	DeleteOrder(ctx context.Context, id int64) error
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
