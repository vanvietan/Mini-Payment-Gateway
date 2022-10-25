package transaction

import (
	"net/http"
	"pg/internal/handler/common"
)

// InitAuthentication check cards information and generate otp
func (h Handler) InitAuthentication(w http.ResponseWriter, r *http.Request) {
	cardInput, orderInput, err := checkValidationAndAmount(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	card, order, errB := h.TxSvc.InitAuthentication(r.Context(), cardInput, orderInput)
	if errB != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJSON(w, http.StatusOK, toGetAInitAuthenticateResponse(card, order))
}
