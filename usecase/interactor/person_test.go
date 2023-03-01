package interactor_test

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"testing"
)

var pi inputport.PersonInputport

func TestPersonInteractor_Store(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		args    inputdata.AddPerson
		want    int // check to last insertID only
		wantErr bool
	}{
		{
			name: "store successfully",
			args: inputdata.AddPerson{
				LoginID:    "login_id5",
				PersonName: "no name",
				Password:   "password",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "duplicate loginID",
			args: inputdata.AddPerson{
				LoginID:    "login_id1",
				PersonName: "no name",
				Password:   "password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pi.Store(ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("personInteractor.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPersonInteractor_Update(t *testing.T) {
	ctx := context.Background()
	type args struct {
		pid int
		p   inputdata.UpdatePerson
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "coverage earn",
			args: args{
				pid: 1,
				p: inputdata.UpdatePerson{
					LoginID:    "new_id",
					PersonName: "new_name",
					Password:   "new_pass",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pi.Update(ctx, tt.args.pid, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("personInteractor.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPersonInteractor_GetByID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		personID   int
		wantPerson model.Person
		wantErr    bool
	}{
		{
			name:     "select successfully",
			personID: 1,
			wantPerson: model.Person{
				PersonID:   1,
				PersonName: "name01",
				FriendCode: 2593519733,
			},
			wantErr: false,
		},
		{
			name:     "[another person]select successfully",
			personID: 4,
			wantPerson: model.Person{
				PersonID:   4,
				PersonName: "p_name04",
				FriendCode: 1050136379,
			},
			wantErr: false,
		},
		{
			name:     "unregistered person select",
			personID: 99,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPerson, err := pi.GetByID(ctx, tt.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("personInteractor.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			type checkList struct {
				PersonID   int
				PersonName string
				FriendCode int
				IsCompare  bool
			}
			checkGot := checkList{
				PersonID:   gotPerson.PersonID,
				PersonName: gotPerson.PersonName,
				FriendCode: gotPerson.FriendCode,
				IsCompare:  gotPerson.IsCompare,
			}
			checkWant := checkList{
				PersonID:   tt.wantPerson.PersonID,
				PersonName: tt.wantPerson.PersonName,
				FriendCode: tt.wantPerson.FriendCode,
				IsCompare:  tt.wantPerson.IsCompare,
			}
			if !reflect.DeepEqual(checkGot, checkWant) {
				t.Errorf("personInteractor.GetByID() = %v, want %v", checkGot, checkWant)
			}
		})
	}
}
