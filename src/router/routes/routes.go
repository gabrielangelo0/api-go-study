package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

func Configurate(r *mux.Router) *mux.Router {
	routes := routesClients

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
