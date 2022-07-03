package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sekareco_srv/domain/infra"
	"sekareco_srv/infra/logger"
	"sekareco_srv/infra/timer"
	"strings"
	"time"
)

var EXPIRED_IN = 1 * time.Hour
var MAX_TOKENS = 100

// RFC 6750 Bearer Token Conform
const (
	REQUEST_HEADER  = "Authorization"
	RESPONSE_HEADER = "WWW-Authenticate"
)

const (
	MESSAGE_OK            = "Bearer realm=\"\""
	MESSAGE_UNAUTHORIZED  = "Bearer realm=\"token_required\""
	MESSAGE_BAD_REQUEST   = "Bearer error=\"invalid_request\""
	MESSAGE_INVALID_TOKEN = "Bearer error=\"invalid_token\""
	MESSAGE_FORBIDDEN     = "Bearer error=\"insufficient_scope\""
)

var Auth infra.Auth

type Access struct {
	token     string
	expiredIn time.Time
}

type AuthManager struct {
	// access token mapping
	// key: personID, value: token
	tokens map[int]*Access
}

func InitAuth() {
	Auth = &AuthManager{
		tokens: make(map[int]*Access, MAX_TOKENS),
	}
}

func (a *AuthManager) GenerateNewToken() string {
	return base64.StdEncoding.EncodeToString([]byte(timer.Timer.NowDatetime()))
}

func (a *AuthManager) AddToken(pid int, token string) {
	a.tokens[pid] = &Access{
		token:     token,
		expiredIn: timer.Timer.Add(EXPIRED_IN),
	}
}

func (a *AuthManager) RevokeToken(pid int) {
	delete(a.tokens, pid)
}

// using middleware
func (a *AuthManager) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := a.getInHeaderToken(r)
		if len(token) == 0 {
			logger.Logger.Warn(fmt.Errorf("%s", "unauthorized"))
			w.Header().Set(RESPONSE_HEADER, MESSAGE_UNAUTHORIZED)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: person ID getting from request parameter
		pid := 1
		if !a.isEnabledToken(pid, token) {
			logger.Logger.Warn(fmt.Errorf("%s", "invalid token"))
			w.Header().Set(RESPONSE_HEADER, MESSAGE_INVALID_TOKEN)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set(RESPONSE_HEADER, MESSAGE_OK)
		next.ServeHTTP(w, r)
	})
}

func (a *AuthManager) getInHeaderToken(r *http.Request) string {
	token := r.Header.Get(REQUEST_HEADER)
	return strings.Trim(strings.Replace(token, "Bearer", "", -1), " ")
}

func (a *AuthManager) isEnabledToken(pid int, token string) bool {
	access, ok := a.tokens[pid]

	// not exist token or token unmatch or token expired
	return !ok || token != access.token || timer.Timer.Before(access.expiredIn)
}
