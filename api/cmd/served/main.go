package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"pg/api/cmd/served/router"
	"pg/api/data"
	cardRepo "pg/api/internal/repository/card"
	orderRepo "pg/api/internal/repository/order"
	transactionRepo "pg/api/internal/repository/transaction"
	cardSvc "pg/api/internal/service/card"
	orderSvc "pg/api/internal/service/order"
	transactionSvc "pg/api/internal/service/transaction"
	"pg/api/internal/util"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	//Create database connection
	dbConn, err := data.GetDatabaseConnection()
	if err != nil {
		log.Fatal("encountered error when create a db connection, error :%v", err)
	}
	//sonyflake
	util.Init()

	orderRepository := orderRepo.New(dbConn)
	orderService := orderSvc.New(orderRepository)
	cardRepository := cardRepo.New(dbConn)
	cardService := cardSvc.New(cardRepository)
	transactionRepository := transactionRepo.New(dbConn)
	transactionService := transactionSvc.New(transactionRepository, cardRepository, orderRepository)

	router.New(r, cardService, transactionService, orderService)

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
