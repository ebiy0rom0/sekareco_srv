package router

import (
	"sekareco_srv/infra/middleware"
	"sekareco_srv/infra/web"
	"sekareco_srv/interface/database"
	"sekareco_srv/interface/handler"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/interactor"

	"github.com/gorilla/mux"
)

// @BasePath /api/v1
func InitRouter(sh infra.SqlHandler, th infra.TxHandler) *mux.Router {
	// create handler rooting
	r := mux.NewRouter()

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

	// rest api root path prefix
	r.PathPrefix("/api/v1")

	// health check end point
	r.HandleFunc("/health", web.HttpHandler(healthHandler.Get).Exec).Methods("GET")

	// account api
	r.HandleFunc("/signup", web.HttpHandler(personHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/signin", web.HttpHandler(authHandler.Post).Exec).Methods("POST")

	// in-app api needs authentication
	iar := r.PathPrefix("/prsk").Subrouter()

	am := middleware.NewAuthMiddleware()
	iar.Use(am.CheckAuth)

	// account api
	iar.HandleFunc("/signout", web.HttpHandler(authHandler.Delete).Exec).Methods("DELETE")

	// person api
	iar.HandleFunc("/persons/{personID}", web.HttpHandler(personHandler.Get).Exec).Methods("GET")
	iar.HandleFunc("/persons/{personID}", web.HttpHandler(personHandler.Put).Exec).Methods("PUT")

	// music api
	iar.HandleFunc("/musics", web.HttpHandler(musicHandler.Get).Exec).Methods("GET")

	// record api
	iar.HandleFunc("/records/{personID}", web.HttpHandler(recordHandler.Get).Exec).Methods("GET")
	iar.HandleFunc("/records/{personID}", web.HttpHandler(recordHandler.Post).Exec).Methods("POST")
	iar.HandleFunc("/records/{personID}/{musicID}", web.HttpHandler(recordHandler.Put).Exec).Methods("PUT")

	return r
}
