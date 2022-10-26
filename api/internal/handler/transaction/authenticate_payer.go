package transaction

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

func (h Handler) AuthenticatePayer(w http.ResponseWriter, r *http.Request) {
	cardID, orderID, err := checkCardIDAndOrderID(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
	}
	trans, errS := h.TxSvc.CreateTransaction(r.Context(), cardID, orderID)
	if errS != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}
	toOTPResponse(w, trans)
	//common.ResponseJSON(w, http.StatusOK, toAuthenticatePayerResponse(w, trans))
}
