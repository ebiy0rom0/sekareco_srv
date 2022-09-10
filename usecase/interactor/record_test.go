package interactor

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/inputdata"
	"testing"
)

func TestRecordInteractor_Store(t *testing.T) {
	type args struct {
		ctx      context.Context
		personID int
		r        inputdata.AddRecord
	}
	tests := []struct {
		name    string
		l       *recordInteractor
		args    args
		want    model.Record
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Store(tt.args.ctx, tt.args.personID, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("recordInteractor.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recordInteractor.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordInteractor_Update(t *testing.T) {
	type args struct {
		ctx      context.Context
		personID int
		musicID  int
		r        inputdata.UpdateRecord
	}
	tests := []struct {
		name    string
		l       *recordInteractor
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Update(tt.args.ctx, tt.args.personID, tt.args.musicID, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("recordInteractor.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecordInteractor_GetByPersonID(t *testing.T) {
	type args struct {
		ctx      context.Context
		personID int
	}
	tests := []struct {
		name        string
		l           *recordInteractor
		args        args
		wantRecords []model.Record
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := tt.l.GetByPersonID(tt.args.ctx, tt.args.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("recordInteractor.GetByPersonID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("recordInteractor.GetByPersonID() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
