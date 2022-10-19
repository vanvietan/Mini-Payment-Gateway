package router

import (
	"github.com/go-chi/chi/v5"
	"pg/internal/handler"
	cardSvc "pg/internal/service/card"
	orderSvc "pg/internal/service/order"
	transactionSvc "pg/internal/service/transaction"
)

// MasterRoute masterRoute
type MasterRoute struct {
	Router             *chi.Mux
	Handler            handler.Handler
	CardService        cardSvc.Service
	TransactionService transactionSvc.Service
}

// New DI
func New(r *chi.Mux, cardService cardSvc.Service, transactionService transactionSvc.Service, orderService orderSvc.Service) {
	newHandler := handler.New(cardService, transactionService, orderService)
	mr := MasterRoute{
		Router:  r,
		Handler: newHandler,
	}
	mr.initRoutes()
}

func (mr MasterRoute) initRoutes() {
	mr.initCardRoutes()
	mr.initTransactionRoutes()
}

func (mr MasterRoute) initCardRoutes() {
	mr.Router.Group(func(r chi.Router) {
		r.Get("/cards", mr.Handler.CardHandler.GetCards)
		r.Get("/cards/{id}", mr.Handler.CardHandler.GetCardByID)
		r.Post("/cards", mr.Handler.CardHandler.AddCard)
		r.Put("/cards", mr.Handler.CardHandler.UpdateCard)
		r.Delete("/cards/{id}", mr.Handler.CardHandler.DeleteCard)
	})
}
func (mr MasterRoute) initTransactionRoutes() {
	mr.Router.Group(func(r chi.Router) {
		r.Post("/tx", mr.Handler.TransHandler.InitAuthentication)
		r.Get("/form", mr.Handler.TransHandler.OTPForm)
		r.Post("/form", mr.Handler.TransHandler.EnterOTP)
	})
}
