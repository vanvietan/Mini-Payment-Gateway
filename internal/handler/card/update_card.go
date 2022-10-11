package card

import (
	"net/http"
	"pg/internal/handler/common"
)

// UpdateCard update a card
func (h Handler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidation(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	cardID, errI := validateIDAndMap(r)
	if errI != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: errI.Error(),
		})
		return
	}
	cardU, errU := h.CardSvc.UpdateCard(r.Context(), reqBody, cardID)
	if errU != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toGetACardResponse(cardU))
}
