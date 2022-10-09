package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"pg/api/data"
)

func main() {
	port := "2000"
	r := chi.NewRouter()

	//Create db connection
	_, err := data.GetDatabaseConnection()
	if err != nil {
		log.Fatal("encountered error when create a db connection, error :%v", err)
	}

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
