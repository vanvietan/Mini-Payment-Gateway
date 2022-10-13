package order

import (
	"context"
	"pg/internal/model"
	"pg/internal/repository/order"
)

// Service contains all order service
type Service interface {
	//CreateOrder create order
	CreateOrder(ctx context.Context, input model.Order) (model.Order, error)

	//GetOrderByID get an order by its id
	GetOrderByID(ctx context.Context, id int64) (model.Order, error)
}
type impl struct {
	orderRepo order.Repository
}

// New DI
func New(orderRepo order.Repository) Service {
	return impl{
		orderRepo: orderRepo,
	}
}
