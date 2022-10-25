package order

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/api/internal/model"
)

func (i impl) GetOrderByID(ctx context.Context, id int64) (model.Order, error) {
	order, err := i.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		if err != nil {
			log.Printf("error when find ID %v ", err)
			return model.Order{}, err
		}
	}
	return order, nil
}
