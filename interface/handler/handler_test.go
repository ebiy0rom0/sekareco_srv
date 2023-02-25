//go:build integration

package handler_test

import (
	"sekareco_srv/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.Initialize()

	// r := httptest.NewRequest()

	m.Run()
}
