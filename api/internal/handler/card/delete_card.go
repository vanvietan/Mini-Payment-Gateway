package card

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

// DeleteCard delete a card
func (h Handler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	cardID, err := validateIDAndMap(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errD := h.CardSvc.DeleteCard(r.Context(), cardID)
	if errD != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}
	common2.ResponseJSON(w, http.StatusOK, toSuccessDelete())
}
