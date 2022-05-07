package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/interface/database"
	logic "sekareco_srv/logic/record"

	"github.com/gorilla/mux"
)

type RecordHandler struct {
	Logic logic.RecordLogic
}

func NewRecordHandler(handler database.SqlHandler) *RecordHandler {
	return &RecordHandler{
		Logic: logic.RecordLogic{
			Repository: &database.RecordRepository{
				Handler: handler,
			},
		},
	}
}

func (handler *RecordHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "record"
	vars["methods"] = "get"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (handler *RecordHandler) Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "record"
	vars["methods"] = "post"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (handler *RecordHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "record"
	vars["methods"] = "put"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
