package interactor_test

import (
	"context"
	"sekareco_srv/domain/infra"
	"sekareco_srv/usecase/inputport"
	"testing"
)

var ai inputport.AuthInputport
var token infra.Token

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
		{
			name: "auth successfully",
			args: args{
				loginID:  "login_id1",
				password: "password",
			},
			wantPersonID: 1,
			wantErr:      false,
		},
		{
			name: "auth successfully(anotherID)",
			args: args{
				loginID:  "login_id3",
				password: "password",
			},
			wantPersonID: 3,
			wantErr:      false,
		},
		{
			name: "not found loginID",
			args: args{
				loginID:  "login_id",
				password: "password",
			},
			wantErr: true,
		},
		{
			name: "unmatch password",
			args: args{
				loginID:  "login_id1",
				password: "passcode",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPersonID, err := ai.CheckAuth(ctx, tt.args.loginID, tt.args.password)
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
	// coverage earn
	token = ai.AddToken(1)
}

func TestAuthInteractor_RevokeToken(t *testing.T) {
	// coverage earn
	ai.RevokeToken(token)
}
