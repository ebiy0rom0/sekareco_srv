package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	stage := flag.String("stage", "dev", "")
	flag.Parse()

	// load env
	if err := infra.LoadEnv(*stage); err != nil {
		log.Fatal(err)
	}

	// logger setup
	if err := infra.InitLogger(); err != nil {
		log.Fatal(err)
	}

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
	cors := middleware.NewCorsConfig()

	srv := &http.Server{
		Handler:      cors.Handler(r),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// wait http request
	go func () {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ctx, stop := context.WithCancel(context.Background())

	// automatically token revoke
	go am.DeleteExpiredToken(ctx)

	// server graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown request error: %+v", err)
	}
}
