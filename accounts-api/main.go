package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ricardomaricato/payment-routine/accounts-api/config"
	"github.com/ricardomaricato/payment-routine/accounts-api/database"
)

func main() {
	config.Load()
	database.Connect()
	fmt.Println("Iniciando a aplicação")
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
