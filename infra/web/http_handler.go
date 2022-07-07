package web

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
)

type HttpHandler func(context.Context, infra.HttpContext)

func (fn HttpHandler) Exec(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	context := NewHttpContext(w, r)
	fn(ctx, context)
}
