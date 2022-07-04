package router

import (
	"sekareco_srv/infra/sql"
	"sekareco_srv/infra/web"
	"sekareco_srv/interface/controller"
	"sekareco_srv/interface/database"
	"sekareco_srv/interface/handler"
	"sekareco_srv/logic"

	"github.com/gorilla/mux"
)

func InitRouter(h *sql.SqlHandler) *mux.Router {
	// create handler rooting
	r := mux.NewRouter()

	authHandler := handler.NewAuthHandler(h)
	musicHandler := handler.NewMusicHandler(h)
	personHandler := handler.NewPersonHandler(
		logic.NewPersonLogic(
			database.NewPersonRepository(h),
			database.NewLoginRepository(h),
		),
	)
	recordHandler := handler.NewRecordHandler(
		logic.NewRecordLogic(
			database.NewRecordRepository(h),
		),
	)

	// account api
	r.HandleFunc("/signup/", web.HttpHandler(personHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/signin/", web.HttpHandler(authHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/signout/", web.HttpHandler(authHandler.Delete).Exec).Methods("DELETE")

	// in-app api needs authentication
	iar := r.PathPrefix("/app").Subrouter()

	ac := controller.NewAuthController(h)
	iar.Use(ac.CheckAuth)

	// person api
	iar.HandleFunc("/person/{personID}/", web.HttpHandler(personHandler.Get).Exec).Methods("GET")
	iar.HandleFunc("/person/{personID}/", web.HttpHandler(personHandler.Put).Exec).Methods("PUT")

	// music api
	iar.HandleFunc("/music/", web.HttpHandler(musicHandler.Get).Exec).Methods("GET")

	// record api
	iar.HandleFunc("/record/{personID}/", web.HttpHandler(recordHandler.Get).Exec).Methods("GET")
	iar.HandleFunc("/record/{personID}/", web.HttpHandler(recordHandler.Post).Exec).Methods("POST")
	iar.HandleFunc("/record/{personID}/{musicID}/", web.HttpHandler(recordHandler.Put).Exec).Methods("PUT")

	return r
}
