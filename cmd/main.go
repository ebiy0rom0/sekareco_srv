package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	loader "sekareco_srv/infra/config"
	"sekareco_srv/infra/logger"
	"sekareco_srv/infra/router"
	"sekareco_srv/infra/sql"
	"sekareco_srv/infra/timer"

	"github.com/rs/cors"
)

func main() {
	// load env
	if err := loader.LoadEnv(".env.development"); err != nil {
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

	// cors setup
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
		},
		AllowedHeaders: []string{
			"Content-Type",
		},
		AllowedMethods: []string{
			"HEAD", "GET", "POST", "PUT", "OPTIONS",
		},
		AllowCredentials: true,
	})

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
