package transaction

import (
	"errors"
	"net/http"
	"pg/internal/handler/common"
	"pg/internal/model"
)

// EnterOTP enter its otp
func (h Handler) EnterOTP(w http.ResponseWriter, r *http.Request) {
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

	card, err := h.TxSvc.InitPayment(r.Context(), trans.ID)
	if err != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toSuccessResponse(card))

}
func toSuccessResponse(card model.Card) PayResponse {
	return PayResponse{
		Message: "Successful Transaction",
		Number:  card.Number,
		Balance: card.Balance,
	}
}

func checkOTP(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	otp := r.Form.Get("otp")
	//transID := r.Form.Get("trans")
	if otp == "" {
		return "", errors.New("invalid OTP")
	}

	//log.Println(transID)
	//i, _ := strconv.ParseInt(transID, 10, 64)
	//log.Println(i)
	//if err != nil {
	//	return "", 0, errors.New("invalid transID")
	//}
	return otp, nil
}
