package infra

import "net/http"

type Auth interface {
	GenerateNewToken() string
	AddToken(int, string)
	RevokeToken(int)
	CheckAuth(http.Handler) http.Handler
}
