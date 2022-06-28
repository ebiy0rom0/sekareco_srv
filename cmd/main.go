package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"sekareco_srv/infra/config"
	"sekareco_srv/infra/logger"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/router"
	"sekareco_srv/infra/sql"
	"sekareco_srv/infra/timer"
)

func main() {
	// load env
	if err := config.LoadEnv(".env.development"); err != nil {
		fmt.Println(err)
	}

	// timer setup
	timer.InitTimer()

	// logger setup
	logger.InitLogger()
	defer logger.DropLogFile()

	// router setup
	if err := router.InitRouter(); err != nil {
		fmt.Println(err)
	}

	// middleware setup
	router.Router.Use(middleware.LoggingAccessLog)

	// cors setup
	c := middleware.InitCors()

	// server setup
	srv := &http.Server{
		Handler:      c.Handler(router.Router),
		Addr:         "0.0.0.0:8000",
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
