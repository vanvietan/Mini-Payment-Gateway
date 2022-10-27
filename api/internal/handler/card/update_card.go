package card

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

// UpdateCard update a card
func (h Handler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidation(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	cardID, errI := validateID(r)
	if errI != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: errI.Error(),
		})
		return
	}
	cardU, errU := h.CardSvc.UpdateCard(r.Context(), reqBody, cardID)
	if errU != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}

	common2.ResponseJSON(w, http.StatusOK, toGetACardResponse(cardU))
}
