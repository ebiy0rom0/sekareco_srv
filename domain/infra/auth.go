package infra

import "net/http"

type Auth interface {
	GenerateNewToken() string
	AddTokens(int, string)
	RemoveTokens(int)
	IsEnableToken(int, string) bool
	CheckAuth(http.Handler) http.Handler
}
