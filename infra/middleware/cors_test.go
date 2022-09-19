package middleware

import (
	"testing"
)

func TestNewCorsConfig(t *testing.T) {
	t.Run("new cors call only", func(t *testing.T) {
		// coverage earn
		var _ = NewCorsConfig()
	})
}
