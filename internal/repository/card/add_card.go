package card

import (
	"context"
	"pg/internal/model"
)

// AddCard add a card
func (i impl) AddCard(ctx context.Context, input model.Card) (model.Card, error) {
	card := model.Card{}
	card.Number = input.Number
	//tx := i.gormDB.Select("transactions.*").Where("number = ? ", card.Number).FirstOrCreate(&card)
	tx := i.gormDB.Create(&input)
	if tx.Error != nil {
		return model.Card{}, tx.Error
	}
	return card, nil
}
