package card

import (
	"context"
	"pg/internal/model"
)

func (i impl) UpdateCard(ctx context.Context, card model.Card) (model.Card, error) {
	tx := i.gormDB.Save(&card)
	if tx.Error != nil {
		return model.Card{}, tx.Error
	}
	return card, nil
}
