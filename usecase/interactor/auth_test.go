//go:build unit
// +build unit

package interactor_test

import (
	"context"
	"sekareco_srv/domain/infra"
	"sekareco_srv/interface/database"
	"sekareco_srv/test"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/interactor"
	"testing"
)

var i inputport.AuthInputport

func TestMain(m *testing.M) {
	test.Setup()

	tx := database.NewTransaction(test.InjectTxHandler())
	l := database.NewLoginRepository(test.InjectSqlHandler())
	i = interactor.NewAuthInteractor(test.InjectTokenManager(), l, tx)

	m.Run()
}

func TestAuthInteractor_CheckAuth(t *testing.T) {
	ctx := context.Background()
	type args struct {
		loginID  string
		password string
	}
	tests := []struct {
		name         string
		args         args
		wantPersonID int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPersonID, err := i.CheckAuth(ctx, tt.args.loginID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authInteractor.CheckAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPersonID != tt.wantPersonID {
				t.Errorf("authInteractor.CheckAuth() = %v, want %v", gotPersonID, tt.wantPersonID)
			}
		})
	}
}

func TestAuthInteractor_AddToken(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i.AddToken(tt.args.id)
		})
	}
}

func TestAuthInteractor_RevokeToken(t *testing.T) {
	type args struct {
		token infra.Token
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i.RevokeToken(tt.args.token)
		})
	}
}
