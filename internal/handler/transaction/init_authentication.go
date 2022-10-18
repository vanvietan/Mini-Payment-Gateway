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

	_, errT := h.TxSvc.CreateTransaction(r.Context(), card.ID, order.ID)
	if errT != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	//tpl := template.Must(template.New("trans").Parse(strconv.FormatInt(trans.ID, 10)))
	//tpl, _ := template.New("trans").Parse(string(trans.ID))
	//tpl.Execute(w, nil)

	http.Redirect(w, r, "/form", http.StatusOK)
	common.ResponseJSON(w, http.StatusOK, toGetGenerateOTPResponse())
}

func toGetGenerateOTPResponse() OTPResponse {
	return OTPResponse{
		Message: "created a transaction",
	}
}
