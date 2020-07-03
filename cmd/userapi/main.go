package main

import (
	"log"
	"net/http"

	"github.com/evseevbl/userapi/internal/app/router"
	"github.com/evseevbl/userapi/internal/app/userapi"
)

func main() {
	api := userapi.NewImplementation()
	srv := router.New(api)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
