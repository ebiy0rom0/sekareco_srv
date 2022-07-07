package middleware

import (
	"fmt"
	"net/http"
	"sekareco_srv/infra"
)

func LoggingAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		infra.Logger.Info(fmt.Errorf("URL:%s, Host:%s, Method:%s", r.URL, r.Host, r.Method))
		next.ServeHTTP(w, r)
	})
}
