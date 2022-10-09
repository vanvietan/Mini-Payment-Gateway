package card

import (
	"context"
	"gorm.io/gorm"
	"pg/internal/model"
)

// CardRepository contains all card repository functions
type CardRepository interface {
	//GetCards get all card
	GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error)
}

type impl struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) CardRepository {
	return impl{
		gormDB: gormDB,
	}
}
