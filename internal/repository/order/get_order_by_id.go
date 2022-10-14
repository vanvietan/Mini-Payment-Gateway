package order

import (
	"context"
	"pg/internal/model"
)

func (i impl) GetOrderByID(ctx context.Context, id int64) (model.Order, error) {
	card := model.Order{}
	tx := i.gormDB.First(&card, id)
	if tx.Error != nil {
		return model.Order{}, tx.Error
	}
	return card, nil
}
