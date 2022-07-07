package database_test

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"
)

func TestLoginRepository_Store(t *testing.T) {
	type args struct {
		ctx context.Context
		l   model.Login
	}
	tests := []struct {
		name    string
		r       database.LoginRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Store(tt.args.ctx, tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("loginRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginRepository_GetByID(t *testing.T) {
	type args struct {
		ctx     context.Context
		loginID string
	}
	tests := []struct {
		name      string
		r         database.LoginRepository
		args      args
		wantLogin model.Login
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogin, err := tt.r.GetByID(tt.args.ctx, tt.args.loginID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogin, tt.wantLogin) {
				t.Errorf("loginRepository.GetByID() = %v, want %v", gotLogin, tt.wantLogin)
			}
		})
	}
}
