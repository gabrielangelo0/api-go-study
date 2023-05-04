package router

import (
	"github.com/gabriel/estudo-api/src/router/routes"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
