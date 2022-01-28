package router

import (
	"github.com/gorilla/mux"
	"github.com/ricardomaricato/payment-routine/accounts-api/router/routes"
)

// GenerateRouter create routing
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
