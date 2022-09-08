package middleware

import (
	"net/http"
	"regexp"
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
			time.Sleep(10 * time.Nanosecond)
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
		if !authMid.isEnabledToken(firstToken) {
			t.Error("not added token")
		}

		// next generate token ensure changes
		time.Sleep(100 * time.Nanosecond)
		newToken := authMid.GenerateNewToken()

		authMid.AddToken(1, newToken)
		// first token is already updated by new token
		if authMid.isEnabledToken(firstToken) {
			t.Error("not deleted old token")
		}
		if !authMid.isEnabledToken(newToken) {
			t.Error("not added new token")
		}

		// new token has been registered in the current thread
		authMid.RevokeToken(newToken)
		if authMid.isEnabledToken(newToken) {
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
	for i := 0; i < 3; i++ {
		t.Run("header token parse check", func(t *testing.T) {
			// set to bearer token
			token := authMid.GenerateNewToken()
			r.Header["Authorization"] = []string{"Bearer " + string(token)}

			if headerToken := authMid.getHeaderToken(r); headerToken != token {
				t.Errorf("failed to header in bearer token. want: %s, got: %s", token, headerToken)
			}
		})
		// next generate token ensure changes
		time.Sleep(10 * time.Nanosecond)
	}
}
