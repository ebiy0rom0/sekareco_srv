package middleware

import (
	"testing"
)

func TestMain(m *testing.M) {
	authMid = NewAuthMiddleware()

	m.Run()
}
