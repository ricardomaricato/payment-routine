package routes

import "net/http"

var routesAccounts = []Route{
	{
		URI:      "/v1/accounts",
		Method:   http.MethodPost,
		Function: nil,
	},
}
