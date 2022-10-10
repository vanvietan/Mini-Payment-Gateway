package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"pg/api/cmd/served/router"
	"pg/api/data"
	cardRepo "pg/internal/repository/card"
	cardSvc "pg/internal/service/card"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	//Create db connection
	dbConn, err := data.GetDatabaseConnection()
	if err != nil {
		log.Fatal("encountered error when create a db connection, error :%v", err)
	}

	cardRepository := cardRepo.New(dbConn)
	cardService := cardSvc.New(cardRepository)
	router.New(r, cardService)

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
