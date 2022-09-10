package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/router"
	"sekareco_srv/infra/sql"

	"github.com/rs/zerolog"
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
	env := flag.String("env", "dev", "")
	flag.Parse()

	// load env
	if err := infra.LoadEnv(*env); err != nil {
		log.Fatal(err)
	}

	// logger setup
	if err := infra.InitLogger(); err != nil {
		log.Fatal(err)
	}
	defer infra.DropLogFile()

	// sql & tx handler setup
	dbPath := os.Getenv("DB_PATH") + os.Getenv("DB_NAME")
	sh, th, err := sql.NewSqlHandler(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// middleware setup
	am := middleware.NewAuthMiddleware()
	l := zerolog.New(os.Stdout)

	// router setup
	r := router.InitRouter(sh, th, am, l)

	// cors setup
	c := middleware.NewCorsConfig()

	srv := &http.Server{
		Handler:      c.Handler(r),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// automatically token revoke
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	am.DeleteExpiredToken(t)

	// wait http request
	log.Fatal(srv.ListenAndServe())
}
