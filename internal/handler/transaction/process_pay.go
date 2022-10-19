package transaction

import (
	"net/http"
	"pg/internal/handler/common"
)

// ProcessPay process payment
func (h Handler) ProcessPay(w http.ResponseWriter, r *http.Request) {
	transID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	card, errH := h.TxSvc.InitPayment(r.Context(), transID)
	if errH != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toSuccessResponse(card))
}
