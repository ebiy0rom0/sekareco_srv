package infra

import "net/http"

type AuthMiddleware interface {
	CheckAuth(http.Handler) http.Handler
	AddToken(int, string)
	RevokeToken(int)
	DeleteExpiredToken()
}
