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

func InitRouter(h infra.SqlHandler) *mux.Router {
	// create handler rooting
	r := mux.NewRouter()

	// FIXME: db connection is dummy
	authHandler := handler.NewAuthHandler(
		interactor.NewAuthInteractor(
			database.NewLoginRepository(h),
			// database.NewTransaction(db),
			nil,
		),
	)
	musicHandler := handler.NewMusicHandler(
		interactor.NewMusicInteractor(
			database.NewMusicRepository(h),
			// database.NewTransaction(db),
			nil,
		),
	)
	personHandler := handler.NewPersonHandler(
		interactor.NewPersonInteractor(
			database.NewPersonRepository(h),
			database.NewLoginRepository(h),
			// database.NewTransaction(db),
			nil,
		),
	)
	recordHandler := handler.NewRecordHandler(
		interactor.NewRecordInteractor(
			database.NewRecordRepository(h),
			// database.NewTransaction(db),
			nil,
		),
	)

	// account api
	r.HandleFunc("/signup", web.HttpHandler(personHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/signin", web.HttpHandler(authHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/signout", web.HttpHandler(authHandler.Delete).Exec).Methods("DELETE")

	// in-app api needs authentication
	iar := r.PathPrefix("/app").Subrouter()

	am := middleware.NewAuthMiddleware()
	iar.Use(am.CheckAuth)

	// person api
	iar.HandleFunc("/person/{personID}", web.HttpHandler(personHandler.Get).Exec).Methods("GET")
	iar.HandleFunc("/person/{personID}", web.HttpHandler(personHandler.Put).Exec).Methods("PUT")

	// music api
	iar.HandleFunc("/music", web.HttpHandler(musicHandler.Get).Exec).Methods("GET")

	// record api
	iar.HandleFunc("/record/{personID}", web.HttpHandler(recordHandler.Get).Exec).Methods("GET")
	iar.HandleFunc("/record/{personID}", web.HttpHandler(recordHandler.Post).Exec).Methods("POST")
	iar.HandleFunc("/record/{personID}/{musicID}", web.HttpHandler(recordHandler.Put).Exec).Methods("PUT")

	return r
}
