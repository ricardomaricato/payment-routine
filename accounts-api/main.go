package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ricardomaricato/payment-routine/accounts-api/config"
	"github.com/ricardomaricato/payment-routine/accounts-api/database"
	"github.com/ricardomaricato/payment-routine/accounts-api/handlers"
	"github.com/ricardomaricato/payment-routine/accounts-api/repositories"
	"github.com/ricardomaricato/payment-routine/accounts-api/services"
)

func main() {
	config.Load()
	db := database.Connect()

	accountHandler := handlers.NewAccountHandler(
		services.NewAccountService(
			repositories.NewAccountRepository(db),
		),
	)

	r := mux.NewRouter()
	r.HandleFunc("/v1/accounts", accountHandler.CreateAccountHandler).Methods("POST")

	// r := router.GenerateRouter()
	fmt.Printf("Listening at the door: %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
