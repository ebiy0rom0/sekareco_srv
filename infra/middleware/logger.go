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
			ctx := context.WithValue(r.Context(), logKey, logger)

			writer := web.NewResponseWriterWrapper(w, r)

			// access tot unauthorized path
			if next == nil {
				next = http.NotFoundHandler()
			}
			next.ServeHTTP(writer, r.WithContext(ctx))
			logger.Info().Object("httpRequest", writer).Send()
		})
	}
}

// TODO: logger into context is anti pattern
func GetLogger(ctx context.Context) zerolog.Logger {
	return ctx.Value(logKey).(zerolog.Logger)
}

// interface implementation check
var _ zerolog.LogObjectMarshaler = &web.ResponseWriterWrapper{}
