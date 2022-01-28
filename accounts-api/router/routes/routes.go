package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

// Config put all routes inside the router
func Config(r *mux.Router) *mux.Router {
	routes := routesAccounts
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
