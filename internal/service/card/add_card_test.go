package card

import (
	"pg/internal/model"
	"testing"
)

func TestAddCard(t *testing.T) {
	type addCard struct {
		mockInput model.Card
		mockResp  model.Card
		mockErr   error
	}
	type arg struct {
		addCard    addCard
		givenInput model.Card
		expRs      model.Card
		expErr     error
	}

}
