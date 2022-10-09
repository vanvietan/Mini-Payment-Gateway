package card

import (
	"context"
	"pg/internal/model"
)

func (i impl) GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error) {
	//TODO implement me
	panic("implement me")
}
