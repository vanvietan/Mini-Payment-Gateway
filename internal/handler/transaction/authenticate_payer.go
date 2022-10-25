package transaction

import "net/http"

func (h Handler) AuthenticatePayer(w http.ResponseWriter, r *http.Request) {

	//_, errX := h.TxSvc.CreateTransaction(r.Context(), card.ID, order.ID)
	//if errX != nil {
	//	common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
	//	return
	//}
}
