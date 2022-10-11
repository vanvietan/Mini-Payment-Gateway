package card

import (
	"net/http"
	"pg/internal/handler/common"
)

// DeleteCard delete a card
func (h Handler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	cardID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errD := h.CardSvc.DeleteCard(r.Context(), cardID)
	if errD != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJSON(w, http.StatusOK, toSuccessDelete())
}
