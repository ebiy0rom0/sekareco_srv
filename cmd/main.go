package main

import (
	"log"
	"net/http"
	"time"

	"sekareco_srv/infra/router"

	"github.com/joho/godotenv"
)

func main() {
	// TODO: meke infra logger?
	// logger setup
	log.SetFlags(log.Lshortfile)

	// load env
	err := godotenv.Load("./../config/.env")
	if err != nil {
		log.Fatal(err)
	}

	// router setup
	err = router.InitRouter()
	if err != nil {
		log.Fatal(err)
	}

	// server setup
	srv := &http.Server{
		Handler:      router.Router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// wait http request
	log.Fatal(srv.ListenAndServe())
}
