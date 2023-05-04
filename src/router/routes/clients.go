package routes

import (
	"github.com/gabriel/estudo-api/src/controllers"
	"net/http"
)

var routesClients = []Route{
	{
		URI:            "/clients",
		Method:         http.MethodPost,
		Function:       controllers.CreateClient,
		Authentication: false,
	},
	{
		URI:            "/clients",
		Method:         http.MethodGet,
		Function:       controllers.ListClients,
		Authentication: false,
	},
	{
		URI:    "/clients/{clientId}",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		Authentication: false,
	},
	{
		URI:    "/clients/{clientId}",
		Method: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		Authentication: false,
	},
	{
		URI:    "/clients/{clientId}",
		Method: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		Authentication: false,
	},
}
