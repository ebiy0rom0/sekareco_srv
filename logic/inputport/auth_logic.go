package inputport

import "net/http"

type AuthLogic interface {
	CheckAuth(string, string) (int, error)
	GenerateNewToken() string
	AddToken(int, string)
	RevokeToken(int)
	GetHeaderToken(*http.Request) string
	IsEnabledToken(int, string) bool
	DeleteExpiredToken()
}
