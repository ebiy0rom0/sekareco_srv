package router

import (
	"net/http"
	"os"
	"sekareco_srv/infra"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/web"
	"sekareco_srv/interface/database"
	"sekareco_srv/interface/handler"
	infra_ "sekareco_srv/interface/infra"
	"sekareco_srv/usecase/interactor"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

// @BasePath /api/v1
func InitRouter(sh infra_.SqlHandler, th infra_.TxHandler) *mux.Router {
	// create handler rooting
	router := mux.NewRouter()

	// middleware setup
	am := middleware.NewAuthMiddleware()
	fp, _ := os.OpenFile(os.Getenv("LOG_PATH")+os.Getenv("INFO_LOG_FILE_NAME"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	logger := zerolog.New(fp)

	tx := database.NewTransaction(th)

	healthHandler := handler.NewHealthHandler()

	tm := middleware.NewTokenManager(am)
	lr := database.NewLoginRepository(sh)
	ai := interactor.NewAuthInteractor(tm, lr, tx)
	authHandler := handler.NewAuthHandler(ai)

	mr := database.NewMusicRepository(sh)
	mi := interactor.NewMusicInteractor(mr, tx)
	musicHandler := handler.NewMusicHandler(mi)

	pr := database.NewPersonRepository(sh)
	pi := interactor.NewPersonInteractor(pr, lr, tx)
	personHandler := handler.NewPersonHandler(pi)

	rr := database.NewRecordRepository(sh)
	ri := interactor.NewRecordInteractor(rr, tx)
	recordHandler := handler.NewRecordHandler(ri)

	// @debug: swagger-UI end point
	router.PathPrefix("/swagger").Handler(infra.SwaggerUI())
	// health check end point
	router.HandleFunc("/health", web.HttpHandler(healthHandler.Get).Exec).Methods(http.MethodGet)

	router.Use(middleware.WithLogger(logger))

	// rest api root path prefix
	appRouter := router.PathPrefix("/api/v1").Subrouter()

	// account api at no auth
	appRouter.HandleFunc("/signup", web.HttpHandler(personHandler.Post).Exec).Methods(http.MethodPost)
	appRouter.HandleFunc("/signin", web.HttpHandler(authHandler.Post).Exec).Methods(http.MethodPost)

	// api needs authentication
	authRouter := appRouter.PathPrefix("").Subrouter()
	authRouter.Use(am.CheckAuth)

	// account api
	authRouter.HandleFunc("/signout", web.HttpHandler(authHandler.Delete).Exec).Methods(http.MethodDelete)

	// person api
	authRouter.HandleFunc("/persons/{person_id:[0-9]+}", web.HttpHandler(personHandler.Get).Exec).Methods(http.MethodGet)
	authRouter.HandleFunc("/persons/{person_id:[0-9]+}", web.HttpHandler(personHandler.Put).Exec).Methods(http.MethodPut)

	// music api
	authRouter.HandleFunc("/musics", web.HttpHandler(musicHandler.Get).Exec).Methods(http.MethodGet)

	// record api
	authRouter.HandleFunc("/records/{person_id:[0-9]+}", web.HttpHandler(recordHandler.Get).Exec).Methods(http.MethodGet)
	authRouter.HandleFunc("/records/{person_id:[0-9]+}", web.HttpHandler(recordHandler.Post).Exec).Methods(http.MethodPost)
	authRouter.HandleFunc("/records/{person_id:[0-9]+}/{music_id:[0-9]+}", web.HttpHandler(recordHandler.Put).Exec).Methods(http.MethodPut)

	go am.DeleteExpiredToken()

	return router
}
