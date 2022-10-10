package router

import (
	"github.com/go-chi/chi/v5"
	"pg/internal/handler"
	cardSvc "pg/internal/service/card"
)

// MasterRoute masterRoute
type MasterRoute struct {
	Router      *chi.Mux
	Handler     handler.Handler
	CardService cardSvc.Service
}

// New DI
func New(r *chi.Mux, cardService cardSvc.Service) {
	newHandler := handler.New(cardService)
	mr := MasterRoute{
		Router:  r,
		Handler: newHandler,
	}
	mr.initRoutes()
}

func (mr MasterRoute) initRoutes() {
	mr.initCardRoutes()
}

func (mr MasterRoute) initCardRoutes() {
	mr.Router.Group(func(r chi.Router) {
		r.Get("/cards", mr.Handler.CardHandler.GetCards)
		r.Get("/cards/{id}", mr.Handler.CardHandler.GetCardByID)
	})
}
