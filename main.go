package main

import (
	"fmt"
	"github.com/gabriel/estudo-api/src/config"
	"github.com/gabriel/estudo-api/src/db"
	"github.com/gabriel/estudo-api/src/router"
	"log"
	"net/http"
)

func main() {
	config.Load()

	db.InitFirebaseApp()

	r := router.Generate()

	fmt.Printf("listening in port: %d\n", config.PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
