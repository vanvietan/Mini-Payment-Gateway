package transaction

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"math"
	"net/http"
	"pg/api/internal/model"
	"strconv"
)

// PayResponse message of successful transaction
type PayResponse struct {
	Message string `json:"message"`
	Number  string `json:"number"`
	Balance int64  `json:"balance"`
}

func validateID(r *http.Request) (int64, error) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, errors.New("id must be a number")
	}
	if ID <= 0 || ID > math.MaxInt64 {
		return 0, errors.New("invalid id")
	}
	return ID, nil
}

func toSuccessResponse(card model.Card) PayResponse {
	return PayResponse{
		Message: "Successful Transaction",
		Number:  card.Number,
		Balance: card.Balance,
	}
}
