package person

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
