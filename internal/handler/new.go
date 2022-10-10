package handler

import (
	"pg/internal/handler/card"
	cardService "pg/internal/service/card"
)

// Handler
type Handler struct {
	CardHandler card.Handler
}

// New
func New(cardSvc cardService.Service) Handler {
	return Handler{
		CardHandler: card.Handler{
			CardSvc: cardSvc,
		},
	}
}
