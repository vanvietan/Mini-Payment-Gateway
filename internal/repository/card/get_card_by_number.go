package card

import (
	"context"
	"pg/internal/model"
)

func (i impl) GetCardByNumber(ctx context.Context, number string) (model.Card, error) {
	card := model.Card{}
	tx := i.gormDB.Where("number = ?", number).First(&card)
	if tx.Error != nil {
		return model.Card{}, tx.Error
	}
	return card, nil
}
