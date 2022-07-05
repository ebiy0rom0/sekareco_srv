package model

import "net/http"

// RFC 6750 Bearer Token Conform
const (
	REQUEST_HEADER  = "Authorization"
	RESPONSE_HEADER = "WWW-Authenticate"
)

const (
	HEADER_DONE          = "Bearer realm=\"\""
	HEADER_UNAUTHORIZED  = "Bearer realm=\"token_required\""
	HEADER_BAD_REQUEST   = "Bearer error=\"invalid_request\""
	HEADER_INVALID_TOKEN = "Bearer error=\"invalid_token\""
	HEADER_FORBIDDEN     = "Bearer error=\"insufficient_scope\""
)

type AuthLogic interface {
	CheckAuth(string, string) (int, error)
	GenerateNewToken() string
	AddToken(int, string)
	RevokeToken(int)
	GetHeaderToken(*http.Request) string
	IsEnabledToken(int, string) bool
	DeleteExpiredToken()
}
