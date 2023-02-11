package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sekareco_srv/domain/infra"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/outputdata"
	"testing"

	"github.com/gorilla/mux"
)

func TestHttpContext_Vars(t *testing.T) {
	route := mux.NewRouter()

	route.HandleFunc("/vars/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		httpCtx := NewHttpContext(w, r)
		want := map[string]string{
			"id": "1234567890",
		}

		if got := httpCtx.Vars(); !reflect.DeepEqual(got, want) {
			t.Errorf("HttpContext.Vars() = %v, want %v", got, want)
		}

		w.WriteHeader(http.StatusProcessing)
	}).Methods("GET")

	t.Run("vars number to GET", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/vars/1234567890", nil)

		route.ServeHTTP(w, r)

		if w.Result().StatusCode != http.StatusProcessing {
			t.Error("Not passing expected test cases.")
		}
	})

	route.HandleFunc("/vars/{id:[0-9]*}/{id2:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		httpCtx := NewHttpContext(w, r)
		want := map[string]string{
			"id":  "123",
			"id2": "456",
		}

		if got := httpCtx.Vars(); !reflect.DeepEqual(got, want) {
			t.Errorf("HttpContext.Vars() = %v, want %v", got, want)
		}

		w.WriteHeader(http.StatusProcessing)
	}).Methods("POST")

	t.Run("number vars to POST", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/vars/123/456", nil)

		route.ServeHTTP(w, r)

		if w.Result().StatusCode != http.StatusProcessing {
			t.Error("Not passing expected test cases.")
		}
	})
}

func TestHttpContext_Decode(t *testing.T) {
	want := model.Person{
		PersonID:   1,
		PersonName: "test",
		FriendCode: 123456789,
		IsCompare:  false,
	}

	j, err := json.Marshal(want)
	if err != nil {
		t.Fatal("failed to json marshal")
	}
	reader := bytes.NewReader(j)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/decode", reader)

	t.Run("decode to struct", func(t *testing.T) {
		httpCtx := NewHttpContext(w, r)
		var got model.Person
		if err := httpCtx.Decode(&got); err != nil {
			t.Errorf("HttpContext.Decode() error=%v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("unmatch decode result: want=%+v but got=%+v", want, got)
		}
	})

	// to coverage earn
	t.Run("failed to unmarshal", func(t *testing.T) {
		httpCtx := NewHttpContext(w, r)
		var got string
		if err := httpCtx.Decode(&got); err == nil {
			t.Errorf("somehow Decode succeeded: %s", got)
		}
	})
}

func TestHttpContext_Response(t *testing.T) {
	t.Run("", func(t *testing.T) {
		w := httptest.NewRecorder()
		out := []outputdata.Music{
			{
				MusicID:   221,
				GroupID:   3,
				MusicName: "虚ろを扇ぐ",
				JacketURL: "jacket_0221.png",
				Level:     []int{5, 11, 17, 23, 27},
				Notes:     []int{320, 455, 677, 908, 1112},
			},
		}

		httpCtx := NewHttpContext(w, nil)
		if err := httpCtx.Response(http.StatusOK, out); err != nil {
			t.Errorf("HttpContext.Response() error = %v", err)
		} else {
			contentType := w.Result().Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("invalid content type: %s", contentType)
			}
			if w.Result().StatusCode != http.StatusOK {
				t.Errorf("invalid status code: %d", w.Result().StatusCode)
			}

			buf := bytes.Buffer{}
			buf.ReadFrom(w.Result().Body)

			var got []outputdata.Music
			if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
				t.Fatal("failed to unmarshal to response body.")
			}
			if !reflect.DeepEqual(out, got) {
				t.Errorf("request body unmatch: want=%+v but got=%+v", out, got)
			}
		}

	})

	// to coverage earn
	t.Run("failed to marshal", func(t *testing.T) {
		w := httptest.NewRecorder()
		httpCtx := NewHttpContext(w, nil)
		if err := httpCtx.Response(http.StatusOK, map[bool]interface{}{true: "error"}); err == nil {
			t.Errorf("somehow marshal succeeded.")
		}
	})
}

func TestHttpContext_MakeError(t *testing.T) {
	// to coverage earn
	message := errors.New("test error message")
	want := infra.HttpError{
		Error: message.Error(),
	}
	t.Run("match error struct", func(t *testing.T) {
		//
		httpCtx := NewHttpContext(nil, nil)
		if got := httpCtx.MakeError(message); !reflect.DeepEqual(*got, want) {
			t.Errorf("want=%+v but got=%+v", want, *got)
		}

	})
}
