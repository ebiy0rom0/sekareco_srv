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

func NewPersonHandler(h database.SqlHandler) *PersonHandler {
	return &PersonHandler{
		logic: logic.PersonLogic{
			Repo: &database.PersonRepo{
				Handler: h,
			},
		},
	}
}

func (c *PersonHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (c *PersonHandler) Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (c *PersonHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
