package main

import (
	"sekareco_srv/handler/music"
	"sekareco_srv/handler/person"
	"sekareco_srv/handler/record"

	"github.com/gorilla/mux"
)

func main() {
	// TODO: cors setup

	// handler rooting
	r := mux.NewRouter()

	// person api
	r.HandleFunc("/person/{personId}/", person.Get).Methods("GET")
	r.HandleFunc("/person/", person.Post).Methods("POST")
	r.HandleFunc("/person/{personId}/", person.Patch).Methods("PATCH")

	// music api
	r.HandleFunc("/music/", music.Get).Methods("GET")

	// record api
	r.HandleFunc("/record/{personId}/", record.Get).Methods("GET")
	r.HandleFunc("/record/{personId}/", record.Post).Methods("POST")
	r.HandleFunc("/record/{personId}/{musicId}/", record.Patch).Methods("PATCH")
}
