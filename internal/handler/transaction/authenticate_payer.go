package transaction

import (
	"net/http"
	"pg/internal/handler/common"
)

// AuthenticatePayer enter its otp
func (h Handler) AuthenticatePayer(w http.ResponseWriter, r *http.Request) {
	inputOTP, err := checkOTP(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	transID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errA := h.TxSvc.AuthenticateTransaction(r.Context(), transID, inputOTP)
	if errA != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, toWrongOTPResponse())
		return
	}

	common.ResponseJSON(w, http.StatusOK, toAuthenticatePayerResponse())

}
