package record

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
