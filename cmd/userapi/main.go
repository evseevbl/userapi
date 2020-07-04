package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // registers `postgres` driver

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/evseevbl/userapi/internal/app/router"
	"github.com/evseevbl/userapi/internal/app/userapi"
	"github.com/evseevbl/userapi/internal/pkg/store/pgstore"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")

	db, err := sqlx.Connect(
		"postgres",
		dsn,
	)

	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot connect to db"))
	}

	store := pgstore.New(db)                // storage
	api := userapi.NewImplementation(store) // userAPI
	srv := router.New(api)

	fmt.Println("userapi started")
	log.Fatal(http.ListenAndServe(":8080", srv))
}
