package order

import (
	"context"
	log "github.com/sirupsen/logrus"
	"pg/internal/model"
	"pg/internal/util"
)

var getNextIDFunc = util.GetNextId

func (i impl) CreateOrder(ctx context.Context, input model.Order) (model.Order, error) {
	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return model.Order{}, err
	}
	input.ID = ID

	order, err := i.orderRepo.CreateOrder(ctx, input)
	if err != nil {
		log.Printf("error when add a card: %+v", input)
		return model.Order{}, err
	}
	return order, nil
}
