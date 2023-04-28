package web

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
)

// A HttpHandler API handler function wrapper.
// If you create a new handler, please create it in the form of HttpHandler.
type HttpHandler func(context.Context, infra.HttpContext) *infra.HttpError

// Exec is a function that registers with the mux/router.
// Wrap the created handler with HttpHandler and register it.
func (fn HttpHandler) Exec(w http.ResponseWriter, r *http.Request) {
	if err := fn(r.Context(), NewHttpContext(w, r)); err != nil {
		context.Response(err.Code, struct{ Error string }{Error: err.Msg})
	}
}
