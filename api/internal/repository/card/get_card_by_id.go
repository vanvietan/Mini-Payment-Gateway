package card

import (
	"context"
	"pg/api/internal/model"
)

func (i impl) GetCardByID(ctx context.Context, cardID int64) (model.Card, error) {
	card := model.Card{}
	tx := i.gormDB.First(&card, cardID)
	if tx.Error != nil {
		return model.Card{}, tx.Error
	}
	return card, nil
}
