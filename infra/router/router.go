package router

import (
	"sekareco_srv/infra/sql"
	"sekareco_srv/infra/web"
	"sekareco_srv/interface/handler"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func InitRouter() (err error) {
	// create handler rooting
	r := mux.NewRouter()

	h, err := sql.NewSqlHandler()
	if err != nil {
		return err
	}

	authHandler := handler.NewAuthHandler(h)
	musicHandler := handler.NewMusicHandler(h)
	personHandler := handler.NewPersonHandler(h)
	recordHandler := handler.NewRecordHandler(h)

	// auth api
	r.HandleFunc("/auth/signIn/", web.HttpHandler(authHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/auth/signOut/", web.HttpHandler(authHandler.Delete).Exec).Methods("DELETE")

	// person api
	r.HandleFunc("/person/{personID}/", web.HttpHandler(personHandler.Get).Exec).Methods("GET")
	r.HandleFunc("/person/", web.HttpHandler(personHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/person/{personID}/", web.HttpHandler(personHandler.Put).Exec).Methods("PUT")

	// music api
	r.HandleFunc("/music/", web.HttpHandler(musicHandler.Get).Exec).Methods("GET")

	// record api
	r.HandleFunc("/record/{personID}/", web.HttpHandler(recordHandler.Get).Exec).Methods("GET")
	r.HandleFunc("/record/{personID}/", web.HttpHandler(recordHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/record/{personID}/{musicID}/", web.HttpHandler(recordHandler.Put).Exec).Methods("PUT")

	Router = r
	return
}
