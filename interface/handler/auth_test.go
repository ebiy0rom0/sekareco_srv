//go:build integration

package handler_test

import (
	"context"
	"sekareco_srv/interface/handler"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
	"testing"
)

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		a inputport.AuthInputport
	}
	tests := []struct {
		name string
		args args
		want handler.AuthHandle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := NewAuthHandler(tt.args.a); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewAuthHandler() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_authHandler_Post(t *testing.T) {
	type args struct {
		ctx context.Context
		hc  infra.HttpContext
	}
	tests := []struct {
		name string
		h    handler.AuthHandle
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.h.Post(tt.args.ctx, tt.args.hc)
		})
	}
}

func Test_authHandler_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		hc  infra.HttpContext
	}
	tests := []struct {
		name string
		h    handler.AuthHandle
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.h.Delete(tt.args.ctx, tt.args.hc)
		})
	}
}
