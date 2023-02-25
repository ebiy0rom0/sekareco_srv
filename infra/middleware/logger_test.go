package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	_ "sekareco_srv/infra"
	"strings"
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

func Test_WithLogger(t *testing.T) {
	buf := bytes.Buffer{}
	logger := zerolog.New(&buf)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://api/v1/musics", nil)
	r.Header.Set("X-Forwarded-For", "127.0.0.1")

	testHandler := WithLogger(logger)

	t.Run("access logger test", func(t *testing.T) {
		testHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// something todo...
			w.WriteHeader(http.StatusOK)
		})).ServeHTTP(w, r)

		var log accessLog
		json.Unmarshal(buf.Bytes(), &log)

		// Latency is a variable value that cannot be matched.
		// Exclude from comparison.
		want := accessLog{
			Level: zerolog.LevelInfoValue,
			HttpRequest: httpRequest{
				CacheHit:      false,
				Latency:       log.HttpRequest.Latency,
				Protocol:      "HTTP/1.1",
				RequestMethod: "GET",
				RemoteIp:      "127.0.0.1",
				RequestUrl:    "http://api/v1/musics",
				Status:        http.StatusOK,
			},
		}

		if !reflect.DeepEqual(log, want) {
			t.Errorf("logger content mismatch. want=%+v, log=%+v", want, log)
		}
	})

	// reset buffer
	body := `{"test":⌘}`
	buf = bytes.Buffer{}

	r = httptest.NewRequest("POST", "http://api/v1/paths", strings.NewReader(body))
	r.Header.Set("X-Forwarded-For", "")
	r.RemoteAddr = "192.168.0.1"

	t.Run("access to non valid paths", func(t *testing.T) {
		// If the path is not valid, next becomes nil.
		testHandler(nil).ServeHTTP(w, r)

		var log accessLog
		json.Unmarshal(buf.Bytes(), &log)

		want := accessLog{
			Level: zerolog.LevelInfoValue,
			HttpRequest: httpRequest{
				Body:          body,
				CacheHit:      false,
				Latency:       log.HttpRequest.Latency,
				Protocol:      "HTTP/1.1",
				RequestMethod: "POST",
				RequestSize:   int64(len(body)),
				RemoteIp:      "192.168.0.1",
				RequestUrl:    "http://api/v1/paths",
				Status:        http.StatusNotFound,
			},
		}

		if !reflect.DeepEqual(log, want) {
			t.Errorf("logger content mismatch. want=%+v, log=%+v", want, log)
		}
	})
}
