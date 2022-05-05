package main

import (
	"log"
	"net/http"
	"sekareco_srv/db"
	"sekareco_srv/handler"
	"time"
)

func main() {
	// db (sqlite3) setup
	db.Init()

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
