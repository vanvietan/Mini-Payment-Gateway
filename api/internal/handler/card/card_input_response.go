package card

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"math"
	"net/http"
	"pg/api/internal/model"
	"regexp"
	"strconv"
	"time"
)

const maxNumber = 999999999999

// ACardInput input from clients
type ACardInput struct {
	Number      string    `json:"number"`
	ExpiredDate time.Time `json:"expired_date"`
	CVV         string    `json:"CVV"`
	UserID      int64     `json:"userID"`
	Balance     int64     `json:"balance"`
}

func (c ACardInput) ValidateAndMap() (model.Card, error) {
	if c.UserID <= 0 || c.UserID > math.MaxInt64 {
		return model.Card{}, errors.New("invalid userID")
	}
	match16 := regexp.MustCompile(`^\d{16}$`)
	match3 := regexp.MustCompile(`^\d{3}$`)
	if !match16.MatchString(c.Number) {
		return model.Card{}, errors.New("invalid number")
	}
	if !match3.MatchString(c.CVV) {
		return model.Card{}, errors.New("invalid CVV")
	}
	if c.ExpiredDate.Equal(time.Now()) || c.ExpiredDate.Before(time.Now()) {
		return model.Card{}, errors.New("invalid expired date")
	}
	if c.Balance <= 0 || c.Balance > maxNumber {
		return model.Card{}, errors.New("invalid balance")
	}

	return model.Card{
		Number:      c.Number,
		ExpiredDate: c.ExpiredDate,
		CVV:         c.CVV,
		Balance:     c.Balance,
		UserID:      c.UserID,
	}, nil
}

// ACardResponse response card to clients
type ACardResponse struct {
	ID          int64     `json:"ID"`
	Number      string    `json:"number"`
	ExpiredDate time.Time `json:"expired_date"`
	CVV         string    `json:"CVV"`
	Balance     int64     `json:"balance"`
	UserID      int64     `json:"userID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func validateIDAndMap(r *http.Request) (int64, error) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, errors.New("id must be a number")
	}
	if ID <= 0 || ID > math.MaxInt64 {
		return 0, errors.New("invalid id")
	}
	return ID, nil
}

func toGetACardResponse(card model.Card) ACardResponse {
	return ACardResponse{
		ID:          card.ID,
		Number:      card.Number,
		ExpiredDate: card.ExpiredDate,
		CVV:         card.CVV,
		Balance:     card.Balance,
		UserID:      card.UserID,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
	}
}
