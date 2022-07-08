package middleware

import (
	"net/http"
	"sekareco_srv/infra"
	"strings"
	"time"
)

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

var EXPIRED_IN = 1 * time.Hour
var MAX_TOKENS = 100

var EXPIRED_TOKEN_DELETE_SPAN = 15 * time.Minute

type personToken struct {
	token     string
	expiredIn time.Time
}

type authMiddleware struct {
	// access token mapping
	// key: personID, value: token
	tokens map[int]*personToken
}

func NewAuthMiddleware() *authMiddleware {
	return &authMiddleware{
		tokens: make(map[int]*personToken, MAX_TOKENS),
	}
}

// using middleware
func (m *authMiddleware) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := m.getHeaderToken(r)
		if len(token) == 0 {
			// infra.Logger.Warn(fmt.Errorf("%s", "unauthorized"))
			w.Header().Set(RESPONSE_HEADER, HEADER_UNAUTHORIZED)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: person ID getting from request parameter
		pid := 1
		if !m.isEnabledToken(pid, token) {
			// infra.Logger.Warn(fmt.Errorf("%s", "invalid token"))
			w.Header().Set(RESPONSE_HEADER, HEADER_INVALID_TOKEN)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set(RESPONSE_HEADER, HEADER_DONE)
		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) AddToken(pid int, token string) {
	m.tokens[pid] = &personToken{
		token:     token,
		expiredIn: infra.Timer.Add(EXPIRED_IN),
	}
}

func (m *authMiddleware) RevokeToken(pid int) {
	delete(m.tokens, pid)
}

// automatically delete the expired token at over time
func (m *authMiddleware) DeleteExpiredToken() {
	// t := time.NewTicker(auth.EXPIRED_TOKEN_DELETE_SPAN)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		<-t.C
		m.deleteExpiredToken()
	}
}

func (m *authMiddleware) getHeaderToken(r *http.Request) string {
	token := r.Header.Get(REQUEST_HEADER)
	return strings.Trim(strings.Replace(token, "Bearer", "", -1), " ")
}

func (m *authMiddleware) isEnabledToken(pid int, token string) bool {
	access, ok := m.tokens[pid]

	// not exist token or token unmatch or token expired
	return !ok || token != access.token || infra.Timer.Before(access.expiredIn)
}

func (l *authMiddleware) deleteExpiredToken() {
	for k, t := range l.tokens {
		if infra.Timer.Before(t.expiredIn) {
			l.RevokeToken(k)
		}
	}
}
