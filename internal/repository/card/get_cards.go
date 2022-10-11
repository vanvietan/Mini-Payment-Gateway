package card

import (
	"context"
	"gorm.io/gorm"
	"pg/internal/model"
)

// GetCards get all cards
func (i impl) GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error) {
	cards := make([]model.Card, limit)
	var tx *gorm.DB
	if lastID == 0 {
		tx = i.gormDB.Model(model.Card{}).Select("cards.*").Limit(limit).Order("created_at desc").Find(&cards)
	} else {
		tx = i.gormDB.Select("cards.*").Where("id < ?", lastID).Limit(limit).Order("created_at desc").Find(&cards)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cards, nil

}
