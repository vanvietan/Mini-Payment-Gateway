package transaction

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

// ProcessPay process payment
func (h Handler) ProcessPay(w http.ResponseWriter, r *http.Request) {
	transID, err := validateID(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	card, errH := h.TxSvc.InitPayment(r.Context(), transID)
	if errH != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}

	common2.ResponseJSON(w, http.StatusOK, toSuccessResponse(card))
}
