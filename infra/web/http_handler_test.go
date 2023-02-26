package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sekareco_srv/interface/infra"
	"testing"
)

func TestHttpHandler_Exec(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/handler_test", nil)
	t.Run("Exec do execute function", func(t *testing.T) {
		HttpHandler(func(ctx context.Context, httpCtx infra.HttpContext) *infra.HttpError {
			httpCtx.Response(http.StatusNotFound, "")
			return nil
		}).Exec(w, r)
	})

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("HttpHandler.Exec() not executed: status=%d", w.Result().StatusCode)
	}

	wantErr := "page not found"
	t.Run("Exec returns HttpError", func(t *testing.T) {
		HttpHandler(func(ctx context.Context, httpCtx infra.HttpContext) *infra.HttpError {
			return &infra.HttpError{Msg: wantErr, Code: http.StatusNotFound}
		}).Exec(w, r)
	})

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Unmatch returns http status: status=%d", w.Result().StatusCode)
	}
	if !reflect.DeepEqual(w.Result().Body, struct{ Error string }{Error: wantErr}) {
		t.Errorf("Different body formats. body=%v", w.Result().Body)
	}
}
