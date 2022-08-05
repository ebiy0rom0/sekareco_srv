package middleware

import (
	"encoding/base64"
	"net/http"
	infra_ "sekareco_srv/domain/infra"
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

type tokenStatus struct {
	personID  int
	expiredIn time.Time
}

type AuthMiddleware struct {
	// access token mapping
	// key: token, value: status
	tokens map[string]*tokenStatus
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		tokens: make(map[string]*tokenStatus, MAX_TOKENS),
	}
}

// using middleware
func WithCheckAuth(m *AuthMiddleware) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := m.getHeaderToken(r)
			if len(token) == 0 {
				w.Header().Set(RESPONSE_HEADER, HEADER_UNAUTHORIZED)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !m.isEnabledToken(token) {
				w.Header().Set(RESPONSE_HEADER, HEADER_INVALID_TOKEN)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// make available for each usecase
			ctx := r.Context()
			ctx = infra_.SetToken(ctx, token)
			r = r.WithContext(ctx)

			w.Header().Set(RESPONSE_HEADER, HEADER_DONE)
			next.ServeHTTP(w, r)
		})
	}
}

func (m *AuthMiddleware) GenerateNewToken() string {
	return base64.StdEncoding.EncodeToString([]byte(infra.Timer.NowDatetime()))
}

func (m *AuthMiddleware) AddToken(pid int, token string) {
	m.tokens[token] = &tokenStatus{
		personID:  pid,
		expiredIn: infra.Timer.Add(EXPIRED_IN),
	}
}

func (m *AuthMiddleware) RevokeToken(token string) {
	delete(m.tokens, token)
}

// automatically delete the expired token at over time
func (m *AuthMiddleware) DeleteExpiredToken(t *time.Ticker) {
	go func() {
		for {
			// ticker wait
			<-t.C

			for token, status := range m.tokens {
				if !infra.Timer.Before(status.expiredIn) {
					m.RevokeToken(token)
				}
			}
		}
	}()
}

func (m *AuthMiddleware) getHeaderToken(r *http.Request) string {
	token := r.Header.Get(REQUEST_HEADER)
	return strings.Trim(strings.Replace(token, "Bearer", "", -1), " ")
}

func (m *AuthMiddleware) isEnabledToken(token string) bool {
	access, ok := m.tokens[token]
	// not exist token or token expired
	return ok && infra.Timer.Before(access.expiredIn)
}
