package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/interface/database"
	logic "sekareco_srv/logic/person"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	logic logic.PersonLogic
}

func NewPersonHandler(handler database.SqlHandler) *PersonHandler {
	return &PersonHandler{
		logic: logic.PersonLogic{
			Repository: &database.PersonRepository{
				Handler: handler,
			},
		},
	}
}

func (handler *PersonHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "get"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (handler *PersonHandler) Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "post"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (handler *PersonHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "put"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
