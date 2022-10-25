package card

import (
	"context"
	"pg/api/internal/model"
)

// AddCard add a card
func (i impl) AddCard(ctx context.Context, card model.Card) (model.Card, error) {
	tx := i.gormDB.Create(&card)
	if tx.Error != nil {
		return model.Card{}, tx.Error
	}
	return card, nil
}
