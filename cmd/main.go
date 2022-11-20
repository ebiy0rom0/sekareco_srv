package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sekareco_srv/infra"
	"sekareco_srv/infra/logger"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/router"
	"sekareco_srv/infra/sql"

	"github.com/gorilla/mux"
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
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	// TODO: Move flag setup to infra package.
	stage := flag.String("stage", "dev", "")
	dbUser := flag.String("dbUser", "", "MySQL user name")
	dbPass := flag.String("dbPass", "", "MySQL password")
	dbHost := flag.String("dbHost", "", "MySQL host address")
	flag.Parse()

	if err := infra.LoadEnv(*stage); err != nil {
		return fmt.Errorf("fail to loading dotenv: %+v", err)
	}

	if err := logger.InitLogger(*stage); err != nil {
		return fmt.Errorf("fail to initialize logger: %+v", err)
	}

	// No MySQL setup until performance impact in production,
	// so sqlite3 connections can be obtained for a while.
	con, err := sql.NewConnection(*dbUser, *dbPass, *dbHost, os.Getenv("DB_NAME"))
	if err != nil {
		return fmt.Errorf("fail connect database: %+v", err)
	}

	// sql & tx handler setup
	sh := sql.NewSqlHandler(con)
	th := sql.NewTxHandler(con)

	// middleware setup
	am := middleware.NewAuthMiddleware()
	l := zerolog.New(os.Stdout)
	r := router.InitRouter(sh, th, am, l)

	cors := middleware.NewCorsConfig()
	srv := &http.Server{
		Handler:      cors.Handler(r),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// wait http request
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	go func() {
		rr := mux.NewRouter()
		rr.HandleFunc("/mainte", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("mainte"))
		}).Methods("GET")
		ns := &http.Server{
			Handler:      cors.Handler(rr),
			Addr:         "0.0.0.0:8001",
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
		}
		if err := ns.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// automatically token revoke
	ctx, stop := context.WithCancel(context.Background())
	go am.DeleteExpiredToken(ctx)

	// server graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}
