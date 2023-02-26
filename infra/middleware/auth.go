package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"sekareco_srv/domain/infra"
	"strconv"
	"strings"
	"time"
)

// RFC 6750 Bearer Token Conform
const (
	REQUEST_HEADER  = "Authorization"
	RESPONSE_HEADER = "WWW-Authenticate"
)

// RFC 2617 Authentication header field pattern
const (
	HEADER_DONE          = "Bearer realm=\"\""
	HEADER_UNAUTHORIZED  = "Bearer realm=\"token_required\""
	HEADER_BAD_REQUEST   = "Bearer error=\"invalid_request\""
	HEADER_INVALID_TOKEN = "Bearer error=\"invalid_token\""
	HEADER_FORBIDDEN     = "Bearer error=\"insufficient_scope\""
)

// Token lifetime: 1 hour after generate
var EXPIRED_IN = 1 * time.Hour

var MAX_TOKENS = 30

// Automatically token delete span: Every 1 second
// = max token life is 1hour and 1 second,
// but it doesn't no have to be strictly 1 hour.
var EXPIRED_TOKEN_DELETE_SPAN = 1 * time.Second

// A sessionInfo is stored token expiration at single person.
type sessionInfo struct {
	personID  int
	expiredIn time.Time
}

// A AuthMiddleware is manages all tokens in this service.
type AuthMiddleware struct {
	// access token mapping
	// key: token, value: status
	session map[infra.Token]sessionInfo
	// 1 personID has only 1 token
	personToToken map[int]infra.Token
}

// NewAuthMiddleware returns AuthMiddleware pointer.
//
// [feature]
// Fixed upper limit of holding 30 tokens,
// because max cost in the process is memory allocate.
// (= max number of connections)
// Use sync.Pool to make changes that minimize memory allocation
// while making the number of tokens holding valiable.
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		session:       make(map[infra.Token]sessionInfo, MAX_TOKENS),
		personToToken: make(map[int]infra.Token, MAX_TOKENS),
	}
}

// WithCheckAuth checks if the user with access is an already authenticated user.
// Register with router middleware for endpoints requiring authentication
// to block access by unauthenticated users.
func (m *AuthMiddleware) WithCheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := m.getHeaderToken(r)
		if len(token) == 0 {
			w.Header().Set(RESPONSE_HEADER, HEADER_UNAUTHORIZED)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !m.isEffectiveToken(token) {
			w.Header().Set(RESPONSE_HEADER, HEADER_INVALID_TOKEN)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// make available for each usecase
		personID, ok := m.getPersonID(token)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx := r.Context()
		ctx = infra.SetToken(ctx, infra.NewSessionInfo(personID, token))
		r = r.WithContext(ctx)

		w.Header().Set(RESPONSE_HEADER, HEADER_DONE)
		next.ServeHTTP(w, r)
	})
}

// GenerateNewToken returns new access token.
//
// [investigation]
// Since access times are used to generate tokens, is there a possibility
// of token conflict when concurrent accesses occur?
func (m *AuthMiddleware) GenerateNewToken() infra.Token {
	nano := infra.Timer.NowTime().UnixNano() / 1e6
	seed := []byte(strconv.FormatInt(nano, 10))
	token := base64.StdEncoding.EncodeToString(seed)
	return infra.Token(token)
}

// AddToken is add the new token into process cache.
// If for some reason the authenticated person adds the token again,
// it will be overwritten with the new token.
func (m *AuthMiddleware) AddToken(pid int, new infra.Token) {
	// delete old token
	old, ok := m.personToToken[pid]
	if ok {
		m.RevokeToken(old)
	}

	m.session[new] = sessionInfo{
		personID:  pid,
		expiredIn: infra.Timer.Add(EXPIRED_IN),
	}
	m.personToToken[pid] = new
}

// RevokeToken is delete the specified token from middleware.
func (m *AuthMiddleware) RevokeToken(token infra.Token) {
	pid := m.session[token].personID
	delete(m.personToToken, pid)
	delete(m.session, token)
}

// DeleteExpiredToken is automatically delete the expired token at over time.
func (m *AuthMiddleware) DeleteExpiredToken(ctx context.Context) {
	t := time.NewTicker(EXPIRED_TOKEN_DELETE_SPAN)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		// ticker wait
		case _, ok := <-t.C:
			if !ok {
				continue
			}
			for token, info := range m.session {
				if !infra.Timer.Before(info.expiredIn) {
					m.RevokeToken(token)
				}
			}
		}
	}
}

// getHeaderToken returns the Bearer token in the request header.
func (m *AuthMiddleware) getHeaderToken(r *http.Request) infra.Token {
	token := r.Header.Get(REQUEST_HEADER)
	return infra.Token(strings.Trim(strings.Replace(token, "Bearer", "", -1), " "))
}

// isEffectiveToken reports weather request token is effective.
func (m *AuthMiddleware) isEffectiveToken(token infra.Token) bool {
	session, ok := m.session[token]
	// not exist token or token expired
	return ok && infra.Timer.Before(session.expiredIn)
}

func (m *AuthMiddleware) getPersonID(token infra.Token) (int, bool) {
	session, ok := m.session[token]
	if !ok {
		return 0, ok
	}
	return session.personID, ok
}
