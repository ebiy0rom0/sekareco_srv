package database_test

import (
	"context"
	"database/sql"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"
)

var loginRepo database.LoginRepository

// this database check is unused transaction
func TestLoginRepository_Store(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		args    model.Login
		wantErr bool
	}{
		{
			name: "insert successfully",
			args: model.Login{
				LoginID:      "login_id4",
				PersonID:     4,
				PasswordHash: "$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC",
			},
			wantErr: false,
		},
		{
			name: "duplicate loginID(primary key)",
			args: model.Login{
				LoginID:      "login_id1",
				PersonID:     5,
				PasswordHash: "$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC",
			},
			wantErr: true,
		},
		{
			name: "duplicate personID(unique key)",
			args: model.Login{
				LoginID:      "login_id5",
				PersonID:     1,
				PasswordHash: "$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loginRepo.Store(ctx, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("loginRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		loginID string
		wantErr bool
	}{
		{name: "select successfully", loginID: "login_id1", wantErr: false},
		{name: "not registered loginID", loginID: "login_id5", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogin, err := loginRepo.GetByID(ctx, tt.loginID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err == sql.ErrNoRows {
				return
			}
			if !reflect.DeepEqual(gotLogin.PersonID, 1) {
				t.Errorf("loginRepository.GetByID() = %s, want %s", gotLogin.LoginID, tt.loginID)
			}
		})
	}
}
