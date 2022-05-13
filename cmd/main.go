package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	loader "sekareco_srv/infra/config"
	logger "sekareco_srv/infra/log"
	"sekareco_srv/infra/router"
	"sekareco_srv/infra/sql"
)

func main() {
	// load env
	if err := loader.LoadEnv(".env.development"); err != nil {
		fmt.Println(err)
	}

	// logger setup
	logger.InitLogger()
	defer func() {
		if err := logger.CleanupLogger(); err != nil {
			fmt.Println(err)
		}
	}()

	// router setup
	if err := router.InitRouter(); err != nil {
		fmt.Println(err)
	}

	// server setup
	srv := &http.Server{
		Handler:      router.Router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	// for debug: drop sqlite3 database file
	defer func() {
		if err := sql.DropDB(); err != nil {
			fmt.Println(err)
		}
	}()

	// wait http request
	log.Fatal(srv.ListenAndServe())
}
