package card

import "pg/internal/service/card"

// Handler handle card calls
type Handler struct {
	CardSvc card.Service
}

// New DI
func New(cardService card.Service) Handler {
	return Handler{CardSvc: cardService}
}
