package infra

import (
	"sekareco_srv/infra/sql"
	"sekareco_srv/infra/web"
	"sekareco_srv/interface/handler"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func Init() (err error) {
	// create handler rooting
	r := mux.NewRouter()

	h, err := sql.NewSqlHandler()
	if err != nil {
		return err
	}

	musicHandler := handler.NewMusicHandler(h)
	personHandler := handler.NewPersonHandler(h)
	recordHandler := handler.NewRecordHandler(h)

	// auth api
	// r.HandleFunc("/auth/", auth.Get).Methods("GET")

	// person api
	r.HandleFunc("/person/{personId}/", web.HttpHandler(personHandler.Get).Exec).Methods("GET")
	r.HandleFunc("/person/", web.HttpHandler(personHandler.Post).Exec).Methods("POST")
	r.HandleFunc("/person/{personId}/", web.HttpHandler(personHandler.Put).Exec).Methods("PUT")

	// music api
	r.HandleFunc("/music/", musicHandler.Get).Methods("GET")

	// record api
	r.HandleFunc("/record/{personId}/", recordHandler.Get).Methods("GET")
	r.HandleFunc("/record/{personId}/", recordHandler.Post).Methods("POST")
	r.HandleFunc("/record/{personId}/{musicId}/", recordHandler.Put).Methods("PUT")

	Router = r
	return
}
