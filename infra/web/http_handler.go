package web

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
)

// A HttpHandler API handler function wrapper.
// If you create a new handler, please create it in the form of HttpHandler.
type HttpHandler func(context.Context, infra.HttpContext)

// Exec is a function that registers with the mux/router.
// Wrap the created handler with HttpHandler and register it.
func (fn HttpHandler) Exec(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	context := NewHttpContext(w, r)
	fn(ctx, context)
}
