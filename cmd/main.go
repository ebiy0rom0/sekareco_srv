package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/router"
	"sekareco_srv/infra/sql"
)

func main() {
	// load env
	if err := infra.LoadEnv(".env.development"); err != nil {
		fmt.Println(err)
	}

	// timer setup
	infra.InitTimer()

	// logger setup
	infra.InitLogger()
	defer infra.DropLogFile()

	// common sql handler setup
	dbPath := os.Getenv("DATABASE_SETUP_PATH")
	h, err := infra.NewSqlHandler(dbPath)
	if err != nil {
		fmt.Println(err)
	}

	// router setup
	r := router.InitRouter(h)

	// common middleware setup
	r.Use(middleware.LoggingAccessLog)

	// cors setup
	c := middleware.InitCors()

	// server setup
	srv := &http.Server{
		Handler:      c.Handler(r),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// for debug: drop sqlite3 database file
	defer func() {
		if err := sql.DropDB(dbPath); err != nil {
			fmt.Println(err)
		}
	}()

	// wait http request
	log.Fatal(srv.ListenAndServe())
}
