package order

import (
	"context"
	"pg/internal/model"
)

func (i impl) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	tx := i.gormDB.Create(&order)
	if tx.Error != nil {
		return model.Order{}, tx.Error
	}
	return order, nil
}
