package web

import (
	"net/http"
	"sekareco_srv/interface/handler"
)

type HttpHandler func(handler.HttpContext)

func (fn HttpHandler) Exec(w http.ResponseWriter, r *http.Request) {
	context := NewHttpContext(w, r)
	fn(context)
}
