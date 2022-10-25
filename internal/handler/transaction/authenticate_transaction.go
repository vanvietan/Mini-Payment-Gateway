package transaction

import (
	"net/http"
	"pg/internal/handler/common"
)

// AuthenticateTransaction enter its otp
func (h Handler) AuthenticateTransaction(w http.ResponseWriter, r *http.Request) {
	inputOTP, err := checkOTP(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	transID, err := validateID(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errA := h.TxSvc.AuthenticateTransaction(r.Context(), transID, inputOTP)
	if errA != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toAuthenticateTransactionResponse())

}
