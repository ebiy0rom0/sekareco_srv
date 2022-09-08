package middleware

import "testing"

func TestMain(m *testing.M) {
	// coverage earn
	NewCorsConfig()

	authMid = NewAuthMiddleware()
	m.Run()
}
