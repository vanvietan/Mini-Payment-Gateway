package card

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

// GetCardByID get a card by id
func (h Handler) GetCardByID(w http.ResponseWriter, r *http.Request) {
	cardID, err := validateID(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	card, errH := h.CardSvc.GetCardByID(r.Context(), cardID)
	if errH != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}

	common2.ResponseJSON(w, http.StatusOK, toGetACardResponse(card))
}
