package validator_test

import (
	"context"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"testing"
)

var pv inputport.PersonValidator

func Test_personValidator_ValidationAdd(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		p       inputdata.AddPerson
		wantErr bool
	}{
		{
			name: "no valid",
			p: inputdata.AddPerson{
				LoginID:    "login_id10",
				Password:   "J}:6q*x(i4;=A#Rb",
				PersonName: "No Name",
			},
			wantErr: false,
		},
		{
			name: "too short loginID",
			p: inputdata.AddPerson{
				LoginID:    "log",
				Password:   "J}:6q*x(i4;=A#Rb",
				PersonName: "No Name",
			},
			wantErr: true,
		},
		{
			name: "too short password",
			p: inputdata.AddPerson{
				LoginID:    "logi", // check to boundary value
				Password:   "jc|Nd<O",
				PersonName: "No Name",
			},
			wantErr: true,
		},
		{
			name: "duplicate loginID",
			p: inputdata.AddPerson{
				LoginID:    "login_id1",
				Password:   "jc|Nd<O*", // check to boundary value
				PersonName: "No Name",
			},
			wantErr: true,
		},
		{
			name: "person name is not entered",
			p: inputdata.AddPerson{
				LoginID:    "logi",
				Password:   "jc|Nd<O*",
				PersonName: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pv.ValidationAdd(ctx, tt.p); (err != nil) != tt.wantErr {
				t.Errorf("what's error?: %+v", err)
				t.Errorf("personValidator.ValidationAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_personValidator_ValidationUpdate(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		p       inputdata.UpdatePerson
		wantErr bool
	}{
		{
			name: "no valid",
			p: inputdata.UpdatePerson{},	// wip
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pv.ValidationUpdate(ctx, tt.p); (err != nil) != tt.wantErr {
				t.Errorf("personValidator.ValidationUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
