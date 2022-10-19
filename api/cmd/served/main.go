package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"pg/api/cmd/served/router"
	"pg/api/data"
	cardRepo "pg/internal/repository/card"
	orderRepo "pg/internal/repository/order"
	transactionRepo "pg/internal/repository/transaction"
	cardSvc "pg/internal/service/card"
	orderSvc "pg/internal/service/order"
	transactionSvc "pg/internal/service/transaction"
	"pg/internal/util"
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
