package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ricardomaricato/payment-routine/accounts-api/config"
	"github.com/ricardomaricato/payment-routine/accounts-api/router"
)

func main() {
	config.Load()

	r := router.GenerateRouter()
	fmt.Printf("Listening at the door: %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
