package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sekareco_srv/domain/infra"
	"sekareco_srv/infra/logger"
	"sekareco_srv/infra/timer"
)

var Auth infra.Auth

type Access struct {
	token string
	ttl   int
}

type AuthManager struct {
	// access token mapping
	// key: personID, value: token
	tokens map[int]*Access
}

func InitAuth() {
	Auth = &AuthManager{
		tokens: make(map[int]*Access),
	}
}

func (_ *AuthManager) GenerateNewToken() string {
	return base64.StdEncoding.EncodeToString([]byte(timer.Timer.NowDatetime()))
}

func (a *AuthManager) AddTokens(pid int, token string) {
	a.tokens[pid] = &Access{
		token: token,
		ttl:   30,
	}
}

func (a *AuthManager) RemoveTokens(pid int) {
	delete(a.tokens, pid)
}

func (a *AuthManager) IsEnableToken(pid int, token string) bool {
	access, ok := a.tokens[pid]

	// not exist token or token unmatch
	return !ok || token != access.token
}

func (a *AuthManager) authenticated() bool {
	// TODO: check
	return false
}

// using middleware
func (a *AuthManager) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.authenticated() {
			next.ServeHTTP(w, r)
		} else {
			logger.Logger.Error(fmt.Errorf("%s", "unauthorized"))
		}
	})
}
