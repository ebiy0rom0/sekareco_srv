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

	_ "sekareco_srv/doc/api"

	"github.com/rs/zerolog"

	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title        sekareco_srv
// @version      1.0.0-beta
// @description  sekareco REST API server.

// @license.name  MIT License
// @license.url   https://github.com/ebiy0rom0/sekareco_srv/blob/develop/LICENSE

// @host      localhost:8000
// @BasePath  /api/v1
// @schemes   http https

// @securityDefinitions.apikey  Bearer Authentication
// @in                          header
// @name                        Authorization
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

	// sql & tx handler setup
	dbPath := os.Getenv("DATABASE_SETUP_PATH")
	sh, th, err := sql.NewSqlHandler(dbPath)
	if err != nil {
		fmt.Println(err)
	}

	// router setup
	r := router.InitRouter(sh, th)
	// TODO: swagger is middleware
	// move to infra dir
	// swagger
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
	))

	// middleware setup
	logger := zerolog.New(os.Stdout)
	r.Use(middleware.WithLogger(logger))

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
