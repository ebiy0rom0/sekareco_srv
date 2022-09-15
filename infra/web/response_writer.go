package web

import (
	"net/http"
	"sekareco_srv/infra"
	"time"

	"github.com/rs/zerolog"
)

// A ResponseWriterWrapper is http.ResponseWriter wrapper.
// Some expanded features are available.
type ResponseWriterWrapper struct {
	status       int
	responseSize int
	writer       http.ResponseWriter
	request      *http.Request
	start        time.Time
}

// NewResponseWriterWrapper returns new ResponseWriterWrapper
// wrapped the argument w and r and have the start time of the process.
func NewResponseWriterWrapper(w http.ResponseWriter, r *http.Request) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{
		writer:  w,
		request: r,
		start:   infra.Timer.NowTime(),
	}
}

// Flush
func (r *ResponseWriterWrapper) Flush() {
	flusher := r.writer.(http.Flusher)
	flusher.Flush()
	r.status = http.StatusOK
}

// Header returns request header map.
// Equivalent to http.ResponseWriter.Header().
func (r *ResponseWriterWrapper) Header() http.Header {
	return r.writer.Header()
}

// Write
func (r *ResponseWriterWrapper) Write(content []byte) (int, error) {
	r.responseSize = len(content)
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.writer.Write(content)
}

// WriterHeader writes status code in response header.
func (r *ResponseWriterWrapper) WriteHeader(code int) {
	r.status = code
	r.writer.WriteHeader(code)
}

// MarshalZerologObject saves structured logs.
func (r *ResponseWriterWrapper) MarshalZerologObject(e *zerolog.Event) {
	e.Str("requestMethod", r.request.Method)
	e.Str("requestUrl", r.request.URL.String())
	e.Int64("requestSize", r.request.ContentLength)
	e.Int("status", r.status)
	e.Str("referer", r.request.Referer())
	e.Str("latency", infra.Timer.Sub(r.start).String())
	e.Bool("cacheHit", r.status == http.StatusNotModified)
	forwarded := r.Header().Get("X-Forwarded-For")
	if forwarded == "" {
		e.Str("remoteIp", forwarded)
	} else {
		e.Str("remoteIp", r.request.RemoteAddr)
	}
	e.Str("protocol", r.request.Proto)
}

var _ http.ResponseWriter = (*ResponseWriterWrapper)(nil)
var _ zerolog.LogObjectMarshaler = (*ResponseWriterWrapper)(nil)

