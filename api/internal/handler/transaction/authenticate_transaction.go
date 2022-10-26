package transaction

import (
	"net/http"
	common2 "pg/api/internal/handler/common"
)

// AuthenticateTransaction enter its otp
func (h Handler) AuthenticateTransaction(w http.ResponseWriter, r *http.Request) {
	inputOTP, err := checkOTP(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	transID, err := validateID(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errA := h.TxSvc.AuthenticateTransaction(r.Context(), transID, inputOTP)
	if errA != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}

	common2.ResponseJSON(w, http.StatusOK, toAuthenticateTransactionResponse())

}
