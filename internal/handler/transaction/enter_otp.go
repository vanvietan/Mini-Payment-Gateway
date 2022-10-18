package transaction

import (
	"errors"
	"net/http"
	"pg/internal/handler/common"
)

// EnterOTP enter its otp
func (h Handler) EnterOTP(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//otp := r.Form.Get("otp")
	inputOTP, err := checkOTP(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	trans, errC := h.TxSvc.FindTransactionByOTP(r.Context(), inputOTP)
	if errC != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	order, errF := h.OrderSvc.GetOrderByID(r.Context(), trans.OrderID)
	if errF != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	card, errD := h.CardSvc.DeductCard(r.Context(), trans.CardID, order.Amount)
	if errD != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	errG := h.TxSvc.DeleteTransaction(r.Context(), trans.ID)
	if errG != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toSuccessResponse(card))

}

func checkOTP(r *http.Request) (string, error) {
	r.ParseForm()
	otp := r.Form.Get("otp")
	//transID := r.Form.Get("trans")
	if otp == "" {
		return "", errors.New("invalid OTP")
	}
	//i, err := strconv.ParseInt(transID, 10, 64)
	//if err != nil {
	//	return "", 0, errors.New("invalid transID")
	//}
	return otp, nil
}
