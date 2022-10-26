package transaction

import (
	"pg/api/internal/service/transaction"
)

// Handler handler
type Handler struct {
	TxSvc transaction.Service
}

// New DI
func New(txService transaction.Service) Handler {
	return Handler{
		TxSvc: txService,
	}
}
