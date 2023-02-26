package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sekareco_srv/domain/infra"
	_ "sekareco_srv/infra"
	"testing"
	"time"
)

var authMid *AuthMiddleware

// generate token format check (base64)
func TestAuthMiddleware_GenerateNewToken(t *testing.T) {
	r := regexp.MustCompile(`^[a-zA-Z0-9]*=+$`)
	// test 10 count
	for i := 0; i < 3; i++ {
		t.Run("base64 format check", func(t *testing.T) {
			token := authMid.GenerateNewToken()
			if !r.Match([]byte(token)) {
				t.Errorf("unmatch base64 format = %s", token)
			}

			// next generate token ensure changes
			time.Sleep(1 * time.Millisecond)
		})
	}
}

// token sequence check
// generate -> add -> regenerate(register token update) -> revoke
func TestAuthMiddleware_TokenSequence(t *testing.T) {
	t.Run("token sequence", func(t *testing.T) {
		// safe method
		firstToken := authMid.GenerateNewToken()

		authMid.AddToken(1, firstToken)
		if !authMid.isEffectiveToken(firstToken) {
			t.Error("not added token")
		}

		// next generate token ensure changes
		time.Sleep(1 * time.Millisecond)
		newToken := authMid.GenerateNewToken()

		authMid.AddToken(1, newToken)
		// first token is already updated by new token
		if authMid.isEffectiveToken(firstToken) {
			t.Error("not deleted old token")
		}
		if !authMid.isEffectiveToken(newToken) {
			t.Error("not added new token")
		}

		// new token has been registered in the current thread
		authMid.RevokeToken(newToken)
		if authMid.isEffectiveToken(newToken) {
			t.Error("not revoked token")
		}
	})
}

// in header token parse check
func TestAuthMiddleware_getHeaderToken(t *testing.T) {
	r := &http.Request{
		Header: map[string][]string{
			"Authorization": {},
		},
	}
	t.Run("header token parse check", func(t *testing.T) {
		// set to bearer token
		token := authMid.GenerateNewToken()
		r.Header["Authorization"] = []string{"Bearer " + string(token)}

		if headerToken := authMid.getHeaderToken(r); headerToken != token {
			t.Errorf("failed to header in bearer token. want: %s, got: %s", token, headerToken)
		}
	})
}

func TestAuthMiddleware_DeleteExpiredToken(t *testing.T) {
	ctx := context.Background()
	go authMid.DeleteExpiredToken(ctx)
}

func TestAuthMiddleware_WithCheckAuth(t *testing.T) {
	t.Run("no authorization header ", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://tests", nil)

		authMid.WithCheckAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// nothing todo
		})).ServeHTTP(w, r)

		code := w.Result().StatusCode
		header := w.Header().Get(RESPONSE_HEADER)
		if code != http.StatusUnauthorized {
			t.Errorf("invalid status code: want=%d but got=%d", http.StatusUnauthorized, code)
		}
		if header != HEADER_UNAUTHORIZED {
			t.Errorf("invalid header message: want=%s, but got=%s", HEADER_UNAUTHORIZED, header)
		}
	})

	t.Run("header token invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://tests", nil)
		r.Header.Set(REQUEST_HEADER, "invalid")

		authMid.WithCheckAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// nothing todo
		})).ServeHTTP(w, r)

		code := w.Result().StatusCode
		header := w.Header().Get(RESPONSE_HEADER)
		if code != http.StatusUnauthorized {
			t.Errorf("invalid status code: want=%d but got=%d", http.StatusUnauthorized, code)
		}
		if header != HEADER_INVALID_TOKEN {
			t.Errorf("invalid header message: want=%s, but got=%s", HEADER_INVALID_TOKEN, header)
		}
	})

	// add token in check middleware
	token := authMid.GenerateNewToken()
	authMid.AddToken(1, token)

	t.Run("authentication successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://tests", nil)
		r.Header.Set(REQUEST_HEADER, string(token))

		authMid.WithCheckAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			got, err := infra.GetToken(ctx)
			if err != nil {
				t.Error(err)
				return
			}

			if got != token {
				t.Errorf("invalid get token: want=%s but got=%s", token, got)
			}
		})).ServeHTTP(w, r)

		header := w.Header().Get(RESPONSE_HEADER)
		if header != HEADER_DONE {
			t.Errorf("invalid header message: want=%s, but got=%s", HEADER_DONE, header)
		}
	})

	t.Run("token expired", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://tests", nil)
		r.Header.Set(REQUEST_HEADER, "invalid")

		authMid.WithCheckAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// nothing todo
		})).ServeHTTP(w, r)

		// TODO: Change time zone to exceed token expired in.

		code := w.Result().StatusCode
		header := w.Header().Get(RESPONSE_HEADER)
		if code != http.StatusUnauthorized {
			t.Errorf("invalid status code: want=%d but got=%d", http.StatusUnauthorized, code)
		}
		if header != HEADER_INVALID_TOKEN {
			t.Errorf("invalid header message: want=%s, but got=%s", HEADER_INVALID_TOKEN, header)
		}
	})
}
