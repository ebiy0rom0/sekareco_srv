package main

import (
	"log"
	"net/http"
	db "sekareco_srv/interface/database"
	"sekareco_srv/interface/handler"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load("./../config/.env")
	if err != nil {
		log.Fatal(err)
	}

	// db (sqlite3) setup
	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	// api server setup
	srv := &http.Server{
		Handler:      handler.Init(),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// wait http request
	log.Fatal(srv.ListenAndServe())
}
