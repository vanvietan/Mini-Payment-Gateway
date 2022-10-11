package card

import (
	"net/http"
	"pg/internal/handler/common"
)

// GetCardByID get a card by id
func (h Handler) GetCardByID(w http.ResponseWriter, r *http.Request) {
	cardID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	card, errH := h.CardSvc.GetCardByID(r.Context(), cardID)
	if errH != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toGetACardResponse(card))
}
