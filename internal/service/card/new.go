package card

import (
	"context"
	"pg/internal/model"
	"pg/internal/repository/card"
)

// Service contains all card service
type Service interface {
	//GetCards get all cards
	GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error)

	//GetCardByID get card by its id
	GetCardByID(ctx context.Context, cardID int64) (model.Card, error)

	//GetCardByNumber get card by number
	GetCardByNumber(ctx context.Context, cardNumber string) (model.Card, error)

	//AddCard add a card
	AddCard(ctx context.Context, input model.Card) (model.Card, error)

	//DeleteCard delete a card
	DeleteCard(ctx context.Context, cardID int64) error

	//UpdateCard  update a card
	UpdateCard(ctx context.Context, input model.Card, cardID int64) (model.Card, error)

	//DeductCard deduct its balance
	DeductCard(ctx context.Context, id int64, amount int64) (model.Card, error)
}

type impl struct {
	cardRepo card.Repository
}

// New DI
func New(cardRepo card.Repository) Service {
	return impl{
		cardRepo: cardRepo,
	}
}
