package infra

import (
	infra "sekareco_srv/infra/sql"
	"sekareco_srv/interface/handler"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	// create handler rooting
	r := mux.NewRouter()

	h, err := infra.NewSqlHandler()
	if err != nil {
		return nil
	}

	personHandler := handler.NewPersonHandler(h)
	// auth api
	// r.HandleFunc("/auth/", auth.Get).Methods("GET")
	// person api
	r.HandleFunc("/person/{personId}/", personHandler.Get).Methods("GET")
	r.HandleFunc("/person/", personHandler.Post).Methods("POST")
	r.HandleFunc("/person/{personId}/", personHandler.Put).Methods("PATCH")

	// music api
	// r.HandleFunc("/music/", music.Get).Methods("GET")

	// record api
	// r.HandleFunc("/record/{personId}/", record.Get).Methods("GET")
	// r.HandleFunc("/record/{personId}/", record.Post).Methods("POST")
	// r.HandleFunc("/record/{personId}/{musicId}/", record.Patch).Methods("PATCH")

	return r
}
