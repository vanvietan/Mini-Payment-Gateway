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

	card, errC := h.CardSvc.AddCard(r.Context(), cardInput)
	if errC != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	order, errO := h.OrderSvc.CreateOrder(r.Context(), orderInput)
	if errO != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	OTP, errT := h.TxSvc.GenerateOTP(r.Context(), card.ID, order.ID)
	if errT != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toGetGenerateOTPResponse(OTP))
}

func toGetGenerateOTPResponse(s string) OTPResponse {
	return OTPResponse{
		Message: "Here is your OTP code",
		OTP:     s,
	}
}
