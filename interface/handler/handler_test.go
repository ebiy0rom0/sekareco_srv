//go:build integration
// +build integration

package handler_test

import (
	"sekareco_srv/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.Setup()

	// r := httptest.NewRequest()

	m.Run()
}
