package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/evseevbl/userapi/internal/app/router"
	"github.com/evseevbl/userapi/internal/app/userapi"
	"github.com/evseevbl/userapi/internal/pkg/store/pgstore"
)

func main() {
	db, err := sqlx.Connect(
		"postgres",
		"host=localhost dbname=userapi user=postgres password=postgres port=5432 sslmode=disable",
	)

	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot open db"))
	}

	store := pgstore.New(db)
	api := userapi.NewImplementation(store)
	srv := router.New(api)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
