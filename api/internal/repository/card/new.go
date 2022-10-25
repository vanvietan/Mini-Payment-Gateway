package card

import (
	"context"
	"gorm.io/gorm"
	"pg/api/internal/model"
)

// Repository contains all card repository functions
type Repository interface {
	//GetCards get all card
	GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error)

	//GetCardByID get cardID
	GetCardByID(ctx context.Context, cardID int64) (model.Card, error)

	//AddCard add a card
	AddCard(ctx context.Context, card model.Card) (model.Card, error)

	//UpdateCard update a card
	UpdateCard(ctx context.Context, card model.Card) (model.Card, error)

	//DeleteCard delete a card
	DeleteCard(ctx context.Context, cardID int64) error

	//GetCardByNumber get card by its number
	GetCardByNumber(ctx context.Context, number string) (model.Card, error)
}

type impl struct {
	gormDB *gorm.DB
}

// New DI
func New(gormDB *gorm.DB) Repository {
	return impl{
		gormDB: gormDB,
	}
}
