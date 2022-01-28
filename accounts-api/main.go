package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ricardomaricato/payment-routine/accounts-api/database"
)

func main() {
	database.ConnectToDatabase()
	fmt.Println("Iniciando a aplicação")
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
