package database_test

import (
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra"
	"sekareco_srv/interface/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestP_Store(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql-mock: %s", err)
	}
	pr := database.NewPersonRepository(
		&infra.SqlHandler{
			Conn: db,
			Tx:   nil,
		},
	)
	type args struct {
		p model.Person
	}
	tests := []struct {
		name         string
		r            model.PersonRepository
		args         args
		wantPersonID int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{name: "name", r: pr, args: args{}, wantPersonID: 1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPersonID, err := tt.r.Store(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersonRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPersonID != tt.wantPersonID {
				t.Errorf("PersonRepository.Store() = %v, want %v", gotPersonID, tt.wantPersonID)
			}
		})
	}
}

func Test_GetByID(t *testing.T) {
	type args struct {
		personID int
	}
	tests := []struct {
		name     string
		r        model.PersonRepository
		args     args
		wantUser model.Person
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := tt.r.GetByID(tt.args.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersonRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("PersonRepository.GetByID() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
