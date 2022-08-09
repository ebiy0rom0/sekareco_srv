package database_test

import (
	"context"
	"database/sql"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"
)

var personRepo database.PersonRepository

// this database check is unused transaction
func TestPersonRepository_Store(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name         string
		argPerson    model.Person
		wantPersonID int
		wantErr      bool
	}{
		{
			name: "insert successfully",
			argPerson: model.Person{
				PersonName: "p_name04",
				FriendCode: 1050136379,
			},
			wantPersonID: 4,
			wantErr:      false,
		},
		{
			name: "[another person]insert successfully",
			argPerson: model.Person{
				PersonName: "p_name05",
				FriendCode: 1050136378,
			},
			wantPersonID: 5,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPersonID, err := personRepo.Store(ctx, tt.argPerson)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersonRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotPersonID != tt.wantPersonID {
				t.Errorf("PersonRepository.Store() = %v, want %v", gotPersonID, tt.wantPersonID)
			}
		})
	}
}

func TestPersonRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		personID int
		wantUser model.Person
		wantErr  bool
	}{
		{
			name:     "select successfully",
			personID: 1,
			wantUser: model.Person{
				PersonID:   1,
				PersonName: "name01",
				FriendCode: 2593519733,
			},
			wantErr: false,
		},
		{
			name:     "[another person]select successfully",
			personID: 4,
			wantUser: model.Person{
				PersonID:   4,
				PersonName: "p_name04",
				FriendCode: 1050136379,
			},
			wantErr: false,
		},
		{
			name:     "un registered person select",
			personID: 6,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := personRepo.GetByID(ctx, tt.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersonRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err == sql.ErrNoRows {
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("PersonRepository.GetByID() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
