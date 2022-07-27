package router

import (
	"os"
	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/web"
	"sekareco_srv/interface/database"
	"sekareco_srv/interface/handler"
	_infra "sekareco_srv/interface/infra"
	"sekareco_srv/usecase/interactor"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

// @BasePath /api/v1
func InitRouter(sh _infra.SqlHandler, th _infra.TxHandler) *mux.Router {
	// create handler rooting
	router := mux.NewRouter()

	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(
		interactor.NewAuthInteractor(
			database.NewLoginRepository(sh),
			database.NewTransaction(th),
		),
	)
	musicHandler := handler.NewMusicHandler(
		interactor.NewMusicInteractor(
			database.NewMusicRepository(sh),
			database.NewTransaction(th),
		),
	)
	personHandler := handler.NewPersonHandler(
		interactor.NewPersonInteractor(
			database.NewPersonRepository(sh),
			database.NewLoginRepository(sh),
			database.NewTransaction(th),
		),
	)
	recordHandler := handler.NewRecordHandler(
		interactor.NewRecordInteractor(
			database.NewRecordRepository(sh),
			database.NewTransaction(th),
		),
	)

	// @debug: swagger-UI end point
	router.PathPrefix("/swagger").Handler(infra.SwaggerUI())
	// health check end point
	router.HandleFunc("/health", web.HttpHandler(healthHandler.Get).Exec).Methods("GET")

	// middleware setup
	logger := zerolog.New(os.Stdout)
	router.Use(middleware.WithLogger(logger))

	// rest api root path prefix
	appRouter := router.PathPrefix("/api/v1").Subrouter()

	// account api
	appRouter.HandleFunc("/signup", web.HttpHandler(personHandler.Post).Exec).Methods("POST")
	appRouter.HandleFunc("/signin", web.HttpHandler(authHandler.Post).Exec).Methods("POST")

	// in-app api needs authentication
	authRouter := appRouter.PathPrefix("").Subrouter()

	am := middleware.NewAuthMiddleware()
	authRouter.Use(am.CheckAuth)

	// account api
	authRouter.HandleFunc("/signout", web.HttpHandler(authHandler.Delete).Exec).Methods("DELETE")

	// person api
	authRouter.HandleFunc("/persons/{person_id:[0-9]+}", web.HttpHandler(personHandler.Get).Exec).Methods("GET")
	authRouter.HandleFunc("/persons/{person_id:[0-9]+}", web.HttpHandler(personHandler.Put).Exec).Methods("PUT")

	// music api
	authRouter.HandleFunc("/musics", web.HttpHandler(musicHandler.Get).Exec).Methods("GET")

	// record api
	authRouter.HandleFunc("/records/{person_id:[0-9]+}", web.HttpHandler(recordHandler.Get).Exec).Methods("GET")
	authRouter.HandleFunc("/records/{person_id:[0-9]+}", web.HttpHandler(recordHandler.Post).Exec).Methods("POST")
	authRouter.HandleFunc("/records/{person_id:[0-9]+}/{music_id:[0-9]+}", web.HttpHandler(recordHandler.Put).Exec).Methods("PUT")

	return router
}
