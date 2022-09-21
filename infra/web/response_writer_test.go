package web

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	_ "sekareco_srv/infra"
	"sekareco_srv/usecase/outputdata"
	"testing"

	"github.com/rs/zerolog"
)

type httpRequest struct {
	Body          string `json:"body"`
	CacheHit      bool   `json:"cacheHit"`
	Latency       string `json:"latency"`
	Protocol      string `json:"protocol"`
	Referer       string `json:"referer"`
	RequestMethod string `json:"requestMethod"`
	RequestSize   int64  `json:"requestSize"`
	RemoteIp      string `json:"remoteIp"`
	RequestUrl    string `json:"requestUrl"`
	Status        int    `json:"status"`
}

type accessLog struct {
	Level       string      `json:"level"`
	HttpRequest httpRequest `json:"httpRequest"`
}

func TestResponseWriterWrapper_Flush(t *testing.T) {
	t.Run("flush call", func(t *testing.T) {
		w := httptest.NewRecorder()
		wrapper := NewResponseWriterWrapper(w, nil)

		wrapper.Flush()
		if !w.Flushed {
			t.Error("Flusher.Flash() is not called.")
		}
	})
}

func TestResponseWriterWrapper_Header(t *testing.T) {
	t.Run("get header", func(t *testing.T) {
		w := httptest.NewRecorder()
		wrapper := NewResponseWriterWrapper(w, nil)

		if got := wrapper.Header(); !reflect.DeepEqual(w.Header(), got) {
			t.Errorf("get header unmatch: want=%+v but got=%+v", w.Header(), got)
		}
	})
}

func TestResponseWriterWrapper_Write(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		w := httptest.NewRecorder()
		wrapper := NewResponseWriterWrapper(w, nil)

		out := outputdata.Record{}
		data, err := json.Marshal(out)
		if err != nil {
			t.Fatal("failed to struct marshal")
		}

		if _, err := wrapper.Write(data); err != nil {
			t.Error("ResponseWriterWrapper.Write() is returns error.")
			return
		}

		var got outputdata.Record
		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Fatal("failed to read response body.")
		}
		if err := json.Unmarshal(resp, &got); err != nil {
			t.Fatal("failed to result body unmarshal.")
		}
		if !reflect.DeepEqual(got, out) {
			t.Errorf("ResponseWriterWrapper.Write() = %+v, want %+v", got, out)
		}
	})
}

func TestResponseWriterWrapper_WriteHeader(t *testing.T) {
	t.Run("status overwrite", func(t *testing.T) {
		w := httptest.NewRecorder()

		wrapper := NewResponseWriterWrapper(w, nil)

		wrapper.WriteHeader(http.StatusNotFound)
		if w.Result().StatusCode != http.StatusNotFound {
			t.Errorf("failed to status overwrite: %d", w.Result().StatusCode)
		}
	})
}

func TestResponseWriterWrapper_MarshalZerologObject(t *testing.T) {
	t.Run("event checks", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/tests", nil)

		wrapper := NewResponseWriterWrapper(nil, r)

		buf := bytes.Buffer{}
		logger := zerolog.New(&buf)

		logger.Debug().Object("httpRequest", wrapper).Send()

		var out accessLog
		if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
			t.Fatal("failed to unmarshal log message.")
		}
		t.Logf("event is %+v", out)
	})
}
