package web

import (
	"net/http"
	"sekareco_srv/interface/infra"
)

type HttpHandler func(infra.HttpContext)

func (fn HttpHandler) Exec(w http.ResponseWriter, r *http.Request) {
	context := NewHttpContext(w, r)
	fn(context)
}
