package transaction

import (
	"net/http"
	"pg/internal/handler/common"
	"pg/internal/model"
)

// EnterOTPPay enterOTP and Pay
func (h Handler) EnterOTPPay(w http.ResponseWriter, r *http.Request) {
	inputOTP, err := checkInputOTP(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	trans, errC := h.TxSvc.CompareOTP(r.Context(), inputOTP)
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

func toSuccessResponse(card model.Card) PayResponse {
	return PayResponse{
		Message: "Successful Transaction",
		Number:  card.Number,
		Balance: card.Balance,
	}
}
