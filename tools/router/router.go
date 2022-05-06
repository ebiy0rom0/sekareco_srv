package tools

import (
	"sekareco_srv/handler/auth"
	"sekareco_srv/handler/music"
	"sekareco_srv/handler/person"
	"sekareco_srv/handler/record"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	// create handler rooting
	r := mux.NewRouter()

	// auth api
	r.HandleFunc("/auth/", auth.Get).Methods("GET")
	// person api
	r.HandleFunc("/person/{personId}/", person.Get).Methods("GET")
	r.HandleFunc("/person/", person.Post).Methods("POST")
	r.HandleFunc("/person/{personId}/", person.Put).Methods("PATCH")

	// music api
	r.HandleFunc("/music/", music.Get).Methods("GET")

	// record api
	r.HandleFunc("/record/{personId}/", record.Get).Methods("GET")
	r.HandleFunc("/record/{personId}/", record.Post).Methods("POST")
	r.HandleFunc("/record/{personId}/{musicId}/", record.Patch).Methods("PATCH")

	return r
}
