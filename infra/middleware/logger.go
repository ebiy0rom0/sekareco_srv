package middleware

import (
	"context"
	"net/http"
	"sekareco_srv/infra/web"

	"github.com/rs/zerolog"
)

type contextKey string

const logKey contextKey = "log"

func WithLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceId := r.Header.Get("Trace-Id")
			if traceId == "" {
				traceId = "hoge"
			}
			logger = logger.With().Str("Trace-Id", traceId).Logger()
			ctx := context.WithValue(r.Context(), logKey, logger)

			writer := web.NewResponseWriterWrapper(w, r)

			next.ServeHTTP(writer, r.WithContext(ctx))
			logger.Info().Object("httpRequest", writer).Send()
		})
	}
}

func GetLogger(ctx context.Context) zerolog.Logger {
	return ctx.Value(logKey).(zerolog.Logger)
}

var _ zerolog.LogObjectMarshaler = &web.ResponseWriterWrapper{}
