package card

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"pg/api/internal/model"
)

func (i impl) DeleteCard(ctx context.Context, cardID int64) error {
	var tx *gorm.DB
	if tx = i.gormDB.Delete(&model.Card{}, cardID); tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("record not found")
	}
	return nil
}
