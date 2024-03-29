package middleware

import (
	"net/http"
	"sekareco_srv/infra/web"

	"github.com/rs/zerolog"
)

// WithLogger logs request information and processing results using third-party tool zerolog.
// For more information on the log contents, please se web.MarshalZerologObject()
func WithLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := web.NewResponseWriterWrapper(w, r)

			// access to unauthorized path
			if next == nil {
				next = http.NotFoundHandler()
			}

			// In case not found handler, access logs are not collected.
			next.ServeHTTP(writer, r)

			// output log
			logger.Info().Object("accessLog", writer).Send()
		})
	}
}
