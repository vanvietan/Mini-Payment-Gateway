package card

import (
	"errors"
	"math"
	"net/http"
	common2 "pg/api/internal/handler/common"
	"strconv"
)

const maxLimit = 20

// GetCards get all cards
func (h Handler) GetCards(w http.ResponseWriter, r *http.Request) {
	limit, lastID, err := validateAndMap(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	cards, errS := h.CardSvc.GetCards(r.Context(), limit, lastID)
	if errS != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}
	resp := dataToResponseArray(cards)
	common2.ResponseJSON(w, http.StatusOK, toGetCardsResponse(resp))
}

func validateAndMap(r *http.Request) (int, int64, error) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		return 0, 0, errors.New("limit must be a number")
	}
	cursor, err := strconv.ParseInt(r.URL.Query().Get("cursor"), 10, 64)
	if limit < 1 || limit > maxLimit {
		return 0, 0, errors.New("invalid limit")
	}
	if cursor < 0 || cursor > math.MaxInt64 {
		return 0, 0, errors.New("invalid cursor")
	}

	return limit, cursor, nil
}
