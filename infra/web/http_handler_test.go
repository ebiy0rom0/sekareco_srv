package web

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sekareco_srv/interface/infra"
	"testing"
)

func TestHttpHandler_Exec(t *testing.T) {
	{
		type message struct {
			Msg string `json:"message"`
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/handler_test", nil)
		wantMsg := "ok"
		t.Run("Exec do execute function", func(t *testing.T) {
			HttpHandler(func(ctx context.Context, httpCtx infra.HttpContext) *infra.HttpError {
				body := message{Msg: wantMsg}
				httpCtx.Response(http.StatusNotFound, body)
				return nil
			}).Exec(w, r)
		})

		if w.Result().StatusCode != http.StatusNotFound {
			t.Errorf("HttpHandler.Exec() not executed: status=%d", w.Result().StatusCode)
		}
		defer w.Result().Body.Close()
		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("failed to read request body. %s", err)
			return
		}
		var body message
		if err := json.Unmarshal(b, &body); err != nil {
			t.Errorf("failed to marshal request body. %s", err)
			return
		}
		if !reflect.DeepEqual(body, message{Msg: wantMsg}) {
			t.Errorf("Different body formats. body=%+v", body)
		}
	}
	{
		type errMessage struct {
			Err string `json:"error"`
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/handler_test", nil)
		wantErr := "page not found"
		t.Run("Exec returns HttpError", func(t *testing.T) {
			HttpHandler(func(ctx context.Context, httpCtx infra.HttpContext) *infra.HttpError {
				return &infra.HttpError{Msg: wantErr, Code: http.StatusNotFound}
			}).Exec(w, r)
		})
		defer w.Result().Body.Close()

		if w.Result().StatusCode != http.StatusNotFound {
			t.Errorf("Unmatch returns http status: status=%d", w.Result().StatusCode)
		}
		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("failed to read request body. %s", err)
			return
		}
		var body errMessage
		if err := json.Unmarshal(b, &body); err != nil {
			t.Errorf("failed to marshal request body. %s", err)
			return
		}
		if !reflect.DeepEqual(body, errMessage{Err: wantErr}) {
			t.Errorf("Different body formats. body=%+v", body)
		}
	}
}
