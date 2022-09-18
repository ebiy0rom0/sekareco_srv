package validator_test

import (
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"testing"
)

var av inputport.AuthValidator

func Test_authValidator_ValidationPost(t *testing.T) {
	tests := []struct {
		name    string
		a 		inputdata.PostAuth
		wantErr bool
	}{
		{
			name: "no valid",
			a: inputdata.PostAuth{
				LoginID:  "login_id",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "loginID is not entered",
			a: inputdata.PostAuth{
				LoginID:  "",
				Password: "password",
			},
			wantErr: true,
		},
		{
			name: "password is not entered",
			a: inputdata.PostAuth{
				LoginID:  "login_id",
				Password: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := av.ValidationPost(tt.a); (err != nil) != tt.wantErr {
				t.Errorf("authValidator.ValidationPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
