package transaction

import (
	"net/http"
	"pg/internal/handler/common"
)

func (h Handler) AuthenticatePayer(w http.ResponseWriter, r *http.Request) {
	cardID, orderID, err := checkCardIDAndOrderID(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
	}
	trans, errS := h.TxSvc.CreateTransaction(r.Context(), cardID, orderID)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	toOTPResponse(w, trans)
	//common.ResponseJSON(w, http.StatusOK, toAuthenticatePayerResponse(w, trans))
}
