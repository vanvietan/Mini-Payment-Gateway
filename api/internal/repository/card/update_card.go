package card

import (
	"context"
	"pg/api/internal/model"
)

// UpdateCard update a card
func (i impl) UpdateCard(ctx context.Context, card model.Card) (model.Card, error) {
	tx := i.gormDB.Save(&card)
	if tx.Error != nil {
		return model.Card{}, tx.Error
	}
	return card, nil
}
