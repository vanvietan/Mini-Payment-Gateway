package card

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"math"
	"net/http"
	"pg/internal/model"
	"strconv"
	"time"
)

// ACardInput input from clients
type ACardInput struct {
	ID          int64     `json:"ID"`
	ExpiredDate time.Time `json:"expired_date"`
	CVV         int16     `json:"CVV"`
	UserID      int64     `json:"userID"`
}

func (i ACardInput) validateAndMap() (model.Card, error) {
	if i.ID <= 0 || i.ID > math.MaxInt64 {
		return model.Card{}, errors.New("invalid ID")
	}
	//TODO
	return model.Card{}, nil

}

// ACardResponse response card to clients
type ACardResponse struct {
	ID          int64              `json:"ID"`
	ExpiredDate time.Time          `json:"expired_date"`
	CVV         int16              `json:"CVV"`
	Balance     int64              `json:"balance"`
	UserID      int64              `json:"userID"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Histories   []model.AuditTrail `json:"histories,omitempty"`
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
		ExpiredDate: card.ExpiredDate,
		CVV:         card.CVV,
		Balance:     card.Balance,
		UserID:      card.UserID,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
		Histories:   card.AuditTrails,
	}
}
