package database_test

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"
)

func TestPersonRepository_Store(t *testing.T) {
	ctx := context.Background()
	type args struct {
		p model.Person
	}
	tests := []struct {
		name         string
		r            database.PersonRepository
		args         args
		wantPersonID int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPersonID, err := tt.r.Store(ctx, tt.args.p)
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

func TestPersonRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	type args struct {
		personID int
	}
	tests := []struct {
		name     string
		r        database.PersonRepository
		args     args
		wantUser model.Person
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := tt.r.GetByID(ctx, tt.args.personID)
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
