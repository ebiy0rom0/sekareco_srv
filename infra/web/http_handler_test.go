package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sekareco_srv/interface/infra"
	"testing"
)

func TestHttpHandler_Exec(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/handler_test", nil)
	t.Run("Exec do execute function", func(t *testing.T) {
		HttpHandler(func(ctx context.Context, httpCtx infra.HttpContext) {
			httpCtx.Response(http.StatusNotFound, "")
		}).Exec(w, r)
	})

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("HttpHandler.Exec() not executed: status=%d", w.Result().StatusCode)
	}
}
