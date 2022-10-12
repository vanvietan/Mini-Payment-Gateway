package transaction

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"pg/internal/model"
)

// DeleteTransaction delete
func (i impl) DeleteTransaction(ctx context.Context, transID int64) error {
	var tx *gorm.DB
	if tx = i.gormDB.Delete(&model.Transaction{}, transID); tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("record not found")
	}
	return nil
}
