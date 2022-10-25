package router

import (
	"github.com/go-chi/chi/v5"
	"pg/api/internal/handler"
	cardSvc "pg/api/internal/service/card"
	orderSvc "pg/api/internal/service/order"
	transactionSvc "pg/api/internal/service/transaction"
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
		r.Post("/initAuthentication", mr.Handler.TransHandler.InitAuthentication)
		r.Post("/authenticatePayer", mr.Handler.TransHandler.AuthenticatePayer)
		r.Post("/authenticateTransaction/{id}", mr.Handler.TransHandler.AuthenticateTransaction)
		r.Put("/processPay/{id}", mr.Handler.TransHandler.ProcessPay)
	})
}
