package transaction

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

// InitAuthentication check cards information and generate otp
func (h Handler) InitAuthentication(w http.ResponseWriter, r *http.Request) {
	cardInput, orderInput, err := checkValidationAndAmount(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	card, order, errB := h.TxSvc.InitAuthentication(r.Context(), cardInput, orderInput)
	if errB != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}
	common2.ResponseJSON(w, http.StatusOK, toGetAInitAuthenticateResponse(card, order))
}
